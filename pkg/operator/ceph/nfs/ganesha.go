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

// Package nfs for NFS ganesha
package nfs

import (
	"fmt"

	cephv1beta1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1beta1"
	"github.com/rook/rook/pkg/clusterd"
	"github.com/rook/rook/pkg/operator/k8sutil"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create the ganesha server
func (c *GaneshaController) createGanesha(n cephv1beta1.NFSGanesha) error {
	if err := validateGanesha(c.context, n); err != nil {
		return err
	}

	logger.Infof("start running ganesha %s", n.Name)

	for i := 0; i < n.Spec.Server.Active; i++ {
		name := k8sutil.IndexToName(i)

		configName, err := c.generateConfig(n, name)
		if err != nil {
			return fmt.Errorf("failed to create config. %+v", err)
		}

		// start the deployment
		deployment := c.makeDeployment(n, name, configName)
		_, err = c.context.Clientset.ExtensionsV1beta1().Deployments(n.Namespace).Create(deployment)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				return fmt.Errorf("failed to create ganesha deployment. %+v", err)
			}
			logger.Infof("ganesha deployment %s already exists", deployment.Name)
		} else {
			logger.Infof("ganesha deployment %s started", deployment.Name)
		}

		// create a service
		err = c.createGaneshaService(n, name)
		if err != nil {
			return fmt.Errorf("failed to create ganesha service. %+v", err)
		}

		if err = c.addServerToDatabase(n, name); err != nil {
			logger.Warningf("Failed to add ganesha server %s to database. It may already be added. %+v", name, err)
		}
	}

	return nil
}

func (c *GaneshaController) addServerToDatabase(n cephv1beta1.NFSGanesha, name string) error {
	nodeID := getGaneshaNodeID(n, name)
	logger.Infof("Adding ganesha %s to grace db", nodeID)
	return c.context.Executor.ExecuteCommand(false, "", "ganesha-rados-grace", "--pool", n.Spec.RADOS.Pool, "--ns", n.Spec.RADOS.Namespace, "add", nodeID)
}

func (c *GaneshaController) removeServerFromDatabase(n cephv1beta1.NFSGanesha, name string) error {
	nodeID := getGaneshaNodeID(n, name)
	logger.Infof("Removing ganesha %s from grace db", nodeID)
	return c.context.Executor.ExecuteCommand(false, "", "ganesha-rados-grace", "--pool", n.Spec.RADOS.Pool, "--ns", n.Spec.RADOS.Namespace, "remove", nodeID)
}

func (c *GaneshaController) generateConfig(n cephv1beta1.NFSGanesha, name string) (string, error) {

	data := map[string]string{
		"config": getGaneshaConfig(n, name),
	}
	configMap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-%s-%s", appName, n.Name, name),
			Namespace: n.Namespace,
			Labels:    getLabels(n, name),
		},
		Data: data,
	}
	if _, err := c.context.Clientset.CoreV1().ConfigMaps(n.Namespace).Create(configMap); err != nil {
		if errors.IsAlreadyExists(err) {
			if _, err := c.context.Clientset.CoreV1().ConfigMaps(n.Namespace).Update(configMap); err != nil {
				return "", fmt.Errorf("failed to update ganesha config. %+v", err)
			}
			return configMap.Name, nil
		}
		return "", fmt.Errorf("failed to create ganesha config. %+v", err)
	}
	return configMap.Name, nil
}

// Delete the ganesha server
func (c *GaneshaController) deleteGanesha(n cephv1beta1.NFSGanesha) error {
	for i := 0; i < n.Spec.Server.Active; i++ {
		name := k8sutil.IndexToName(i)

		// Remove from grace db
		if err := c.removeServerFromDatabase(n, name); err != nil {
			logger.Warningf("failed to remove server %s from grace db. %+v", name, err)
		}

		// Delete the mds deployment
		k8sutil.DeleteDeployment(c.context.Clientset, n.Namespace, instanceName(n, name))

		// Delete the ganesha service
		options := &metav1.DeleteOptions{}
		err := c.context.Clientset.CoreV1().Services(n.Namespace).Delete(instanceName(n, name), options)
		if err != nil && !errors.IsNotFound(err) {
			logger.Warningf("failed to delete ganesha service. %+v", err)
		}
	}

	return nil
}

func instanceName(n cephv1beta1.NFSGanesha, name string) string {
	return fmt.Sprintf("%s-%s-%s", appName, n.Name, name)
}

func validateGanesha(context *clusterd.Context, n cephv1beta1.NFSGanesha) error {
	// core properties
	if n.Name == "" {
		return fmt.Errorf("missing name")
	}
	if n.Namespace == "" {
		return fmt.Errorf("missing namespace")
	}

	// Client recovery properties
	if n.Spec.RADOS.Pool == "" {
		return fmt.Errorf("missing RADOS.pool")
	}
	if n.Spec.RADOS.Namespace == "" {
		return fmt.Errorf("missing RADOS.namespace")
	}

	// Ganesha server properties
	if n.Spec.Server.Active == 0 {
		return fmt.Errorf("at least one active server required")
	}

	return nil
}
