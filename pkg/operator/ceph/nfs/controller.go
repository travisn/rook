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

// Package nfs to manage a NFS Ganesha server
package nfs

import (
	"reflect"

	"github.com/coreos/pkg/capnslog"
	opkit "github.com/rook/operator-kit"
	cephv1beta1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1beta1"
	"github.com/rook/rook/pkg/clusterd"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

const (
	customResourceName       = "nfsganesha"
	customResourceNamePlural = "nfsganeshas"
)

var logger = capnslog.NewPackageLogger("github.com/rook/rook", "op-nfs")

// FilesystemResource represents the file system custom resource
var NFSGaneshaResource = opkit.CustomResource{
	Name:    customResourceName,
	Plural:  customResourceNamePlural,
	Group:   cephv1beta1.CustomResourceGroup,
	Version: cephv1beta1.Version,
	Scope:   apiextensionsv1beta1.NamespaceScoped,
	Kind:    reflect.TypeOf(cephv1beta1.NFSGanesha{}).Name(),
}

// NFSGaneshaController represents a controller for NFS custom resources
type GaneshaController struct {
	context     *clusterd.Context
	rookImage   string
	cephVersion cephv1beta1.CephVersionSpec
	hostNetwork bool
	ownerRef    metav1.OwnerReference
}

// NewNFSGaneshaController create controller for watching NFS custom resources created
func NewGaneshaController(context *clusterd.Context, rookImage string, cephVersion cephv1beta1.CephVersionSpec, hostNetwork bool, ownerRef metav1.OwnerReference) *GaneshaController {
	return &GaneshaController{
		context:     context,
		rookImage:   rookImage,
		cephVersion: cephVersion,
		hostNetwork: hostNetwork,
		ownerRef:    ownerRef,
	}
}

// StartWatch watches for instances of Filesystem custom resources and acts on them
func (c *GaneshaController) StartWatch(namespace string, stopCh chan struct{}) error {

	resourceHandlerFuncs := cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onAdd,
		UpdateFunc: c.onUpdate,
		DeleteFunc: c.onDelete,
	}

	logger.Infof("start watching filesystem resource in namespace %s", namespace)
	watcher := opkit.NewWatcher(NFSGaneshaResource, namespace, resourceHandlerFuncs, c.context.RookClientset.CephV1beta1().RESTClient())
	go watcher.Watch(&cephv1beta1.NFSGanesha{}, stopCh)

	return nil
}

func (c *GaneshaController) onAdd(obj interface{}) {
	nfsGanesha := obj.(*cephv1beta1.NFSGanesha).DeepCopy()

	err := c.createGanesha(*nfsGanesha)
	if err != nil {
		logger.Errorf("failed to create NFS Ganesha %s. %+v", nfsGanesha.Name, err)
	}
}

func (c *GaneshaController) onUpdate(oldObj, newObj interface{}) {
	oldNFS := oldObj.(*cephv1beta1.NFSGanesha).DeepCopy()
	newNFS := newObj.(*cephv1beta1.NFSGanesha).DeepCopy()

	if !nfsGaneshaChanged(oldNFS.Spec, newNFS.Spec) {
		logger.Debugf("nfs ganesha %s not updated", newNFS.Name)
		return
	}

	logger.Infof("TODO: Update the ganesha server from %d to %d active count", oldNFS.Spec.Server.Active, newNFS.Spec.Server.Active)
}

func (c *GaneshaController) onDelete(obj interface{}) {
	nfsGanesha := obj.(*cephv1beta1.NFSGanesha).DeepCopy()

	err := c.deleteGanesha(*nfsGanesha)
	if err != nil {
		logger.Errorf("failed to delete file system %s. %+v", nfsGanesha.Name, err)
	}
}

func nfsGaneshaChanged(oldNFS, newNFS cephv1beta1.NFSGaneshaSpec) bool {
	if oldNFS.Server.Active != newNFS.Server.Active {
		return true
	}
	return false
}
