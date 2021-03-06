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
	v1alpha1 "github.com/domoinc/kube-valet/pkg/apis/assignments/v1alpha1"
	scheme "github.com/domoinc/kube-valet/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeAssignmentGroupsGetter has a method to return a NodeAssignmentGroupInterface.
// A group's client should implement this interface.
type NodeAssignmentGroupsGetter interface {
	NodeAssignmentGroups() NodeAssignmentGroupInterface
}

// NodeAssignmentGroupInterface has methods to work with NodeAssignmentGroup resources.
type NodeAssignmentGroupInterface interface {
	Create(*v1alpha1.NodeAssignmentGroup) (*v1alpha1.NodeAssignmentGroup, error)
	Update(*v1alpha1.NodeAssignmentGroup) (*v1alpha1.NodeAssignmentGroup, error)
	UpdateStatus(*v1alpha1.NodeAssignmentGroup) (*v1alpha1.NodeAssignmentGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NodeAssignmentGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.NodeAssignmentGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeAssignmentGroup, err error)
	NodeAssignmentGroupExpansion
}

// nodeAssignmentGroups implements NodeAssignmentGroupInterface
type nodeAssignmentGroups struct {
	client rest.Interface
}

// newNodeAssignmentGroups returns a NodeAssignmentGroups
func newNodeAssignmentGroups(c *AssignmentsV1alpha1Client) *nodeAssignmentGroups {
	return &nodeAssignmentGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the nodeAssignmentGroup, and returns the corresponding nodeAssignmentGroup object, and an error if there is any.
func (c *nodeAssignmentGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.NodeAssignmentGroup, err error) {
	result = &v1alpha1.NodeAssignmentGroup{}
	err = c.client.Get().
		Resource("nodeassignmentgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeAssignmentGroups that match those selectors.
func (c *nodeAssignmentGroups) List(opts v1.ListOptions) (result *v1alpha1.NodeAssignmentGroupList, err error) {
	result = &v1alpha1.NodeAssignmentGroupList{}
	err = c.client.Get().
		Resource("nodeassignmentgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeAssignmentGroups.
func (c *nodeAssignmentGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("nodeassignmentgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a nodeAssignmentGroup and creates it.  Returns the server's representation of the nodeAssignmentGroup, and an error, if there is any.
func (c *nodeAssignmentGroups) Create(nodeAssignmentGroup *v1alpha1.NodeAssignmentGroup) (result *v1alpha1.NodeAssignmentGroup, err error) {
	result = &v1alpha1.NodeAssignmentGroup{}
	err = c.client.Post().
		Resource("nodeassignmentgroups").
		Body(nodeAssignmentGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nodeAssignmentGroup and updates it. Returns the server's representation of the nodeAssignmentGroup, and an error, if there is any.
func (c *nodeAssignmentGroups) Update(nodeAssignmentGroup *v1alpha1.NodeAssignmentGroup) (result *v1alpha1.NodeAssignmentGroup, err error) {
	result = &v1alpha1.NodeAssignmentGroup{}
	err = c.client.Put().
		Resource("nodeassignmentgroups").
		Name(nodeAssignmentGroup.Name).
		Body(nodeAssignmentGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nodeAssignmentGroups) UpdateStatus(nodeAssignmentGroup *v1alpha1.NodeAssignmentGroup) (result *v1alpha1.NodeAssignmentGroup, err error) {
	result = &v1alpha1.NodeAssignmentGroup{}
	err = c.client.Put().
		Resource("nodeassignmentgroups").
		Name(nodeAssignmentGroup.Name).
		SubResource("status").
		Body(nodeAssignmentGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the nodeAssignmentGroup and deletes it. Returns an error if one occurs.
func (c *nodeAssignmentGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nodeassignmentgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeAssignmentGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("nodeassignmentgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nodeAssignmentGroup.
func (c *nodeAssignmentGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeAssignmentGroup, err error) {
	result = &v1alpha1.NodeAssignmentGroup{}
	err = c.client.Patch(pt).
		Resource("nodeassignmentgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
