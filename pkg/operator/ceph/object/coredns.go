/*
Copyright 2018 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package object

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/rook/rook/pkg/operator/ceph/controller"
	"github.com/rook/rook/pkg/operator/k8sutil"
	apps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	coreDNSImage        = "k8s.gcr.io/coredns:1.7.0"
	coreDNSResourceName = "rook-ceph-rgw-coredns"
)

const dnsConfigTemplate = `
@ 3600 IN SOA ns.data.local. admin.data.local. (
  2021041600 ; Serial
  3600       ; Refresh
  600        ; Retry
  604800     ; Expire
  600 )      ; Negative Cache TTL
mcg    60 IN CNAME s3.${NAMESPACE}.svc.cluster.local.
*.mcg  60 IN CNAME s3.${NAMESPACE}.svc.cluster.local.
s3     60 IN CNAME ${OBJECT_STORE_NAME}.${NAMESPACE}.svc.cluster.local.
*.s3   60 IN CNAME ${OBJECT_STORE_NAME}.${NAMESPACE}.svc.cluster.local.
`
const dnsCorefileTemplate = `
.:5353 {
	errors
	health {
		lameduck 20s
	}
	ready
	file /etc/coredns/s3config data.local in-addr.arpa ip6.arpa {
		reload 30s
	}
	prometheus %s
	forward . /etc/resolv.conf {
		policy sequential
	}
	cache 900
	reload
}`

func (c *clusterConfig) reconcileCoreDNS(clusterIP string) error {
	if !c.store.Spec.Gateway.EnableCoreDNS {
		logger.Debug("skipping config of coredns")
		return nil
	}

	if err := c.createCoreDNSService(); err != nil {
		return errors.Wrapf(err, "failed to create coredns svc")
	}

	if err := c.createCoreDNSConfigMap(); err != nil {
		return errors.Wrap(err, "failed to create coredns configmap")
	}

	d, err := c.createCoreDNSDeployment()
	if err != nil {
		return errors.Wrap(err, "failed to create coredns deployment")
	}
	if _, err := k8sutil.CreateOrUpdateDeployment(c.context.Clientset, d); err != nil {
		return errors.Wrap(err, "failed to create coredns deployment")
	}
	return nil
}

func (c *clusterConfig) createCoreDNSService() error {

	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      coreDNSResourceName,
			Namespace: c.store.Namespace,
			Labels:    c.getCoreDNSLabels(),
		},
		Spec: v1.ServiceSpec{
			Selector: c.getCoreDNSLabels(),
		},
	}
	if c.clusterSpec.Network.IsHost() {
		svc.Spec.ClusterIP = v1.ClusterIPNone
	}

	//destPort := c.generateLiveProbePort()

	addPort(svc, "udp", 5353, 5353)
	_, err := k8sutil.CreateOrUpdateService(c.context.Clientset, c.store.Namespace, svc)
	return err
}

func (c *clusterConfig) getCoreDNSLabels() map[string]string {
	return map[string]string{
		"coredns":           "rook",
		"rook_object_store": c.store.Name,
	}
}

func (c *clusterConfig) createCoreDNSDeployment() (*apps.Deployment, error) {
	pod, err := c.makeCoreDNSPodSpec()
	if err != nil {
		return nil, err
	}
	replicas := int32(1)
	d := &apps.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      coreDNSResourceName,
			Namespace: c.store.Namespace,
			Labels:    c.getCoreDNSLabels(),
		},
		Spec: apps.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: c.getCoreDNSLabels(),
			},
			Template: pod,
			Replicas: &replicas,
			Strategy: apps.DeploymentStrategy{
				Type: apps.RecreateDeploymentStrategyType,
			},
		},
	}
	k8sutil.AddRookVersionLabelToDeployment(d)
	controller.AddCephVersionLabelToDeployment(c.clusterInfo.CephVersion, d)
	//c.store.Spec.Gateway.Annotations.ApplyToObjectMeta(&d.ObjectMeta)
	//c.store.Spec.Gateway.Labels.ApplyToObjectMeta(&d.ObjectMeta)

	return d, nil
}

func (c *clusterConfig) makeCoreDNSPodSpec() (v1.PodTemplateSpec, error) {
	podSpec := v1.PodSpec{
		Containers:    []v1.Container{c.makeCoreDNSContainer()},
		RestartPolicy: v1.RestartPolicyAlways,
		Volumes:       nil, // TODO,
		HostNetwork:   c.clusterSpec.Network.IsHost(),
		//PriorityClassName: c.store.Spec.Gateway.PriorityClassName,
	}

	// Replace default unreachable node toleration
	k8sutil.AddUnreachableNodeToleration(&podSpec)

	configVolume := v1.Volume{
		Name: "config-volume",
		VolumeSource: v1.VolumeSource{
			ConfigMap: &v1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: coreDNSResourceName}},
		},
	}
	podSpec.Volumes = append(podSpec.Volumes, configVolume)
	//c.store.Spec.Gateway.Placement.ApplyToPodSpec(&podSpec)

	podTemplateSpec := v1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Name:   coreDNSResourceName,
			Labels: c.getCoreDNSLabels(),
		},
		Spec: podSpec,
	}
	//c.store.Spec.Gateway.Annotations.ApplyToObjectMeta(&podTemplateSpec.ObjectMeta)
	//c.store.Spec.Gateway.Labels.ApplyToObjectMeta(&podTemplateSpec.ObjectMeta)

	if c.clusterSpec.Network.IsHost() {
		podTemplateSpec.Spec.DNSPolicy = v1.DNSClusterFirstWithHostNet
	} else if c.clusterSpec.Network.IsMultus() {
		if err := k8sutil.ApplyMultus(c.clusterSpec.Network, &podTemplateSpec.ObjectMeta); err != nil {
			return podTemplateSpec, err
		}
	}

	return podTemplateSpec, nil
}

func (c *clusterConfig) makeCoreDNSContainer() v1.Container {
	// start the rgw daemon in the foreground
	container := v1.Container{
		Name:  "coredns",
		Image: coreDNSImage,
		Args:  []string{"-conf", "/etc/coredns/Corefile"},
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "config-volume",
				MountPath: "/etc/coredns",
				ReadOnly:  true,
			},
		},
		Ports: []v1.ContainerPort{
			{Name: "dns", ContainerPort: 5353, Protocol: v1.ProtocolUDP},
			{Name: "dns-tcp", ContainerPort: 5353, Protocol: v1.ProtocolTCP},
		},
		//Resources:       c.store.Spec.Gateway.Resources,
		//LivenessProbe:   c.generateLiveProbe(),
		SecurityContext: controller.PodSecurityContext(),
	}

	return container
}

func (c *clusterConfig) createCoreDNSConfigMap() error {
	dnsConfig := strings.ReplaceAll(dnsConfigTemplate, "${NAMESPACE}", c.store.Namespace)
	dnsConfig = strings.ReplaceAll(dnsConfig, "${OBJECT_STORE_NAME}", c.store.Name)

	// TODO: Is this the real prometheus endpoint?
	corefile := fmt.Sprintf(dnsCorefileTemplate, "127.0.0.1:9153")

	configMap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      coreDNSResourceName,
			Namespace: c.store.Namespace,
		},
		Data: map[string]string{
			"s3config": dnsConfig,
			"Corefile": corefile,
		},
	}

	// Set owner reference
	err := controllerutil.SetControllerReference(c.store, configMap, c.client.Scheme())
	if err != nil {
		return errors.Wrapf(err, "failed to set owner reference for coredns configmap %q", configMap.Name)
	}

	ctx := context.TODO()
	if _, err := c.context.Clientset.CoreV1().ConfigMaps(c.store.Namespace).Create(ctx, configMap, metav1.CreateOptions{}); err != nil {
		if !kerrors.IsAlreadyExists(err) {
			return errors.Wrap(err, "failed to create ganesha config map")
		}

		logger.Debugf("updating config map %q that already exists", configMap.Name)
		if _, err = c.context.Clientset.CoreV1().ConfigMaps(c.store.Namespace).Update(ctx, configMap, metav1.UpdateOptions{}); err != nil {
			return errors.Wrap(err, "failed to update ganesha config map")
		}
	}

	return nil
}
