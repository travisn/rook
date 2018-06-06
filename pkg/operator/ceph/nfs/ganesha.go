/*
Copyright 2016 The Rook Authors. All rights reserved.

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

// Package nfs for NFS ganesha
package nfs

import (
	"fmt"

	cephv1alpha1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1alpha1"
	"github.com/rook/rook/pkg/clusterd"
	opmon "github.com/rook/rook/pkg/operator/ceph/cluster/mon"
	"github.com/rook/rook/pkg/operator/k8sutil"
	"k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	appName     = "rook-ceph-ganesha"
	ganeshaPort = 2049
)

// Create the ganesha server
func (c *GaneshaController) createGanesha(n cephv1alpha1.NFSGanesha) error {
	if err := validateGanesha(c.context, n); err != nil {
		return err
	}

	logger.Infof("start running ganesha %s", n.Name)

	// start the deployment
	deployment := c.makeDeployment(n)
	_, err := c.context.Clientset.ExtensionsV1beta1().Deployments(n.Namespace).Create(deployment)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			return fmt.Errorf("failed to create mds deployment. %+v", err)
		}
		logger.Infof("ganesha deployment %s already exists", deployment.Name)
	} else {
		logger.Infof("ganesha deployment %s started", deployment.Name)
	}

	// create a service
	err = c.createGaneshaService(n)
	if err != nil {
		return fmt.Errorf("failed to create ganesha service. %+v", err)
	}

	return nil
}

func (c *GaneshaController) createGaneshaService(n cephv1alpha1.NFSGanesha) error {
	labels := getLabels(n)
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:            instanceName(n),
			Namespace:       n.Namespace,
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{c.ownerRef},
		},
		Spec: v1.ServiceSpec{
			Selector: labels,
			Ports: []v1.ServicePort{
				{
					Name:       "nfs",
					Port:       ganeshaPort,
					TargetPort: intstr.FromInt(int(ganeshaPort)),
					Protocol:   v1.ProtocolTCP,
				},
			},
		},
	}
	if c.hostNetwork {
		svc.Spec.ClusterIP = v1.ClusterIPNone
	}

	svc, err := c.context.Clientset.CoreV1().Services(n.Namespace).Create(svc)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			return fmt.Errorf("failed to create ganesha service. %+v", err)
		}
		logger.Infof("ganesha service already created")
		return nil
	}

	logger.Infof("ganesha service running at %s:%d", svc.Spec.ClusterIP, ganeshaPort)
	return nil
}

// Delete the file system
func (c *GaneshaController) deleteGanesha(n cephv1alpha1.NFSGanesha) error {
	// Delete the mds deployment
	k8sutil.DeleteDeployment(c.context.Clientset, n.Namespace, instanceName(n))

	// Delete the ganesha service
	options := &metav1.DeleteOptions{}
	err := c.context.Clientset.CoreV1().Services(n.Namespace).Delete(instanceName(n), options)
	if err != nil && !errors.IsNotFound(err) {
		logger.Warningf("failed to delete ganesha service. %+v", err)
	}

	return nil
}

func instanceName(n cephv1alpha1.NFSGanesha) string {
	return fmt.Sprintf("%s-%s", appName, n.Name)
}

func (c *GaneshaController) makeDeployment(n cephv1alpha1.NFSGanesha) *extensions.Deployment {
	deployment := &extensions.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:            instanceName(n),
			Namespace:       n.Namespace,
			OwnerReferences: []metav1.OwnerReference{c.ownerRef},
		},
	}

	podSpec := v1.PodSpec{
		Containers:    []v1.Container{c.ganeshaContainer(n)},
		RestartPolicy: v1.RestartPolicyAlways,
		Volumes: []v1.Volume{
			{Name: k8sutil.DataDirVolume, VolumeSource: v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}}},
			k8sutil.ConfigOverrideVolume(),
		},
		HostNetwork: c.hostNetwork,
	}
	if c.hostNetwork {
		podSpec.DNSPolicy = v1.DNSClusterFirstWithHostNet
	}
	n.Spec.Server.Placement.ApplyToPodSpec(&podSpec)

	podTemplateSpec := v1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Name:        instanceName(n),
			Labels:      getLabels(n),
			Annotations: map[string]string{},
		},
		Spec: podSpec,
	}

	replicas := int32(n.Spec.Server.Active)
	deployment.Spec = extensions.DeploymentSpec{Template: podTemplateSpec, Replicas: &replicas}
	return deployment
}

func (c *GaneshaController) ganeshaContainer(n cephv1alpha1.NFSGanesha) v1.Container {

	return v1.Container{
		Args: []string{
			"ceph",
			"ganesha",
		},
		Name:  instanceName(n),
		Image: c.rookImage,
		VolumeMounts: []v1.VolumeMount{
			{Name: k8sutil.DataDirVolume, MountPath: k8sutil.DataDir},
			k8sutil.ConfigOverrideMount(),
		},
		Env: []v1.EnvVar{
			{Name: "ROOK_POD_NAME", ValueFrom: &v1.EnvVarSource{FieldRef: &v1.ObjectFieldSelector{FieldPath: "metadata.name"}}},
			{Name: "ROOK_NFS_EXPORT_POOL", Value: n.Spec.Export.Pool},
			{Name: "ROOK_NFS_EXPORT_OBJECT", Value: n.Spec.Export.Object},
			opmon.ClusterNameEnvVar(n.Namespace),
			opmon.EndpointEnvVar(),
			opmon.AdminSecretEnvVar(),
			k8sutil.PodIPEnvVar(k8sutil.PrivateIPEnvVar),
			k8sutil.PodIPEnvVar(k8sutil.PublicIPEnvVar),
			k8sutil.ConfigOverrideEnvVar(),
		},
		Resources: n.Spec.Server.Resources,
	}
}

func getLabels(n cephv1alpha1.NFSGanesha) map[string]string {
	return map[string]string{
		k8sutil.AppAttr:     appName,
		k8sutil.ClusterAttr: n.Namespace,
		"nfs_ganesha":       n.Name,
	}
}

func validateGanesha(context *clusterd.Context, n cephv1alpha1.NFSGanesha) error {
	if n.Name == "" {
		return fmt.Errorf("missing name")
	}
	if n.Namespace == "" {
		return fmt.Errorf("missing namespace")
	}
	if n.Spec.Server.Active == 0 {
		return fmt.Errorf("at least one active server required")
	}

	err := verifyExportExists(context, n.Spec.Export)
	if err != nil {
		return fmt.Errorf("failed to check for export existence. %+v", err)
	}

	return nil
}

func verifyExportExists(context *clusterd.Context, export cephv1alpha1.GaneshaExportSpec) error {
	// TODO: Check if the pool and the RADOS object exist with the exports
	return nil
}
