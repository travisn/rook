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

package fake

import (
	v1alpha1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNFSGaneshas implements NFSGaneshaInterface
type FakeNFSGaneshas struct {
	Fake *FakeCephV1alpha1
	ns   string
}

var nfsganeshasResource = schema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1alpha1", Resource: "nfsganeshas"}

var nfsganeshasKind = schema.GroupVersionKind{Group: "ceph.rook.io", Version: "v1alpha1", Kind: "NFSGanesha"}

// Get takes name of the nFSGanesha, and returns the corresponding nFSGanesha object, and an error if there is any.
func (c *FakeNFSGaneshas) Get(name string, options v1.GetOptions) (result *v1alpha1.NFSGanesha, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nfsganeshasResource, c.ns, name), &v1alpha1.NFSGanesha{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NFSGanesha), err
}

// List takes label and field selectors, and returns the list of NFSGaneshas that match those selectors.
func (c *FakeNFSGaneshas) List(opts v1.ListOptions) (result *v1alpha1.NFSGaneshaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nfsganeshasResource, nfsganeshasKind, c.ns, opts), &v1alpha1.NFSGaneshaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NFSGaneshaList{}
	for _, item := range obj.(*v1alpha1.NFSGaneshaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nFSGaneshas.
func (c *FakeNFSGaneshas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nfsganeshasResource, c.ns, opts))

}

// Create takes the representation of a nFSGanesha and creates it.  Returns the server's representation of the nFSGanesha, and an error, if there is any.
func (c *FakeNFSGaneshas) Create(nFSGanesha *v1alpha1.NFSGanesha) (result *v1alpha1.NFSGanesha, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nfsganeshasResource, c.ns, nFSGanesha), &v1alpha1.NFSGanesha{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NFSGanesha), err
}

// Update takes the representation of a nFSGanesha and updates it. Returns the server's representation of the nFSGanesha, and an error, if there is any.
func (c *FakeNFSGaneshas) Update(nFSGanesha *v1alpha1.NFSGanesha) (result *v1alpha1.NFSGanesha, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nfsganeshasResource, c.ns, nFSGanesha), &v1alpha1.NFSGanesha{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NFSGanesha), err
}

// Delete takes name of the nFSGanesha and deletes it. Returns an error if one occurs.
func (c *FakeNFSGaneshas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nfsganeshasResource, c.ns, name), &v1alpha1.NFSGanesha{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNFSGaneshas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nfsganeshasResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NFSGaneshaList{})
	return err
}

// Patch applies the patch and returns the patched nFSGanesha.
func (c *FakeNFSGaneshas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NFSGanesha, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nfsganeshasResource, c.ns, name, data, subresources...), &v1alpha1.NFSGanesha{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NFSGanesha), err
}
