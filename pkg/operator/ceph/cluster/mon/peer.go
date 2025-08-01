/*
Copyright 2025 The Rook Authors. All rights reserved.

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

package mon

import (
	"github.com/pkg/errors"
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	"github.com/rook/rook/pkg/operator/ceph/controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Cluster) ConfigureClusterPeerSecret() error {
	createKey := true
	if c.spec.Security.CephX.RBDMirrorPeer.CreateMirrorKey != nil && *c.spec.Security.CephX.RBDMirrorPeer.CreateMirrorKey == false {
		createKey = false
	}

	if !createKey {
		logger.Debug("deleting the peer secret if it exists")
		if err := controller.DeleteBootstrapPeerSecret(c.context, c.ClusterInfo); err != nil {
			logger.Errorf("failed to delete bootstrap peer secret: %v", err)
		}
		return nil
	}

	if c.context.Client == nil {
		// TODO: Set the client in the mon failover tests
		return nil
	}
	logger.Info("creating or updating the cluster peer secret")
	if _, err := controller.CreateBootstrapPeerSecret(c.context, c.ClusterInfo, &cephv1.CephCluster{ObjectMeta: metav1.ObjectMeta{Name: c.ClusterInfo.NamespacedName().Name, Namespace: c.Namespace}}, c.ownerInfo); err != nil {
		return errors.Wrap(err, "failed to create cluster rbd bootstrap peer token")
	}
	return nil
}
