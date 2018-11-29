/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1alpha1"
	scheme "github.com/rook/rook/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NFSGaneshasGetter has a method to return a NFSGaneshaInterface.
// A group's client should implement this interface.
type NFSGaneshasGetter interface {
	NFSGaneshas(namespace string) NFSGaneshaInterface
}

// NFSGaneshaInterface has methods to work with NFSGanesha resources.
type NFSGaneshaInterface interface {
	Create(*v1alpha1.NFSGanesha) (*v1alpha1.NFSGanesha, error)
	Update(*v1alpha1.NFSGanesha) (*v1alpha1.NFSGanesha, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NFSGanesha, error)
	List(opts v1.ListOptions) (*v1alpha1.NFSGaneshaList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NFSGanesha, err error)
	NFSGaneshaExpansion
}

// nFSGaneshas implements NFSGaneshaInterface
type nFSGaneshas struct {
	client rest.Interface
	ns     string
}

// newNFSGaneshas returns a NFSGaneshas
func newNFSGaneshas(c *CephV1alpha1Client, namespace string) *nFSGaneshas {
	return &nFSGaneshas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nFSGanesha, and returns the corresponding nFSGanesha object, and an error if there is any.
func (c *nFSGaneshas) Get(name string, options v1.GetOptions) (result *v1alpha1.NFSGanesha, err error) {
	result = &v1alpha1.NFSGanesha{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nfsganeshas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NFSGaneshas that match those selectors.
func (c *nFSGaneshas) List(opts v1.ListOptions) (result *v1alpha1.NFSGaneshaList, err error) {
	result = &v1alpha1.NFSGaneshaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nfsganeshas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nFSGaneshas.
func (c *nFSGaneshas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nfsganeshas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a nFSGanesha and creates it.  Returns the server's representation of the nFSGanesha, and an error, if there is any.
func (c *nFSGaneshas) Create(nFSGanesha *v1alpha1.NFSGanesha) (result *v1alpha1.NFSGanesha, err error) {
	result = &v1alpha1.NFSGanesha{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nfsganeshas").
		Body(nFSGanesha).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nFSGanesha and updates it. Returns the server's representation of the nFSGanesha, and an error, if there is any.
func (c *nFSGaneshas) Update(nFSGanesha *v1alpha1.NFSGanesha) (result *v1alpha1.NFSGanesha, err error) {
	result = &v1alpha1.NFSGanesha{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nfsganeshas").
		Name(nFSGanesha.Name).
		Body(nFSGanesha).
		Do().
		Into(result)
	return
}

// Delete takes name of the nFSGanesha and deletes it. Returns an error if one occurs.
func (c *nFSGaneshas) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nfsganeshas").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nFSGaneshas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nfsganeshas").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nFSGanesha.
func (c *nFSGaneshas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NFSGanesha, err error) {
	result = &v1alpha1.NFSGanesha{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nfsganeshas").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
