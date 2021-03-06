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
	v1alpha1 "github.com/domoinc/kube-valet/pkg/apis/assignments/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterPodAssignmentRules implements ClusterPodAssignmentRuleInterface
type FakeClusterPodAssignmentRules struct {
	Fake *FakeAssignmentsV1alpha1
}

var clusterpodassignmentrulesResource = schema.GroupVersionResource{Group: "assignments.kube-valet.io", Version: "v1alpha1", Resource: "clusterpodassignmentrules"}

var clusterpodassignmentrulesKind = schema.GroupVersionKind{Group: "assignments.kube-valet.io", Version: "v1alpha1", Kind: "ClusterPodAssignmentRule"}

// Get takes name of the clusterPodAssignmentRule, and returns the corresponding clusterPodAssignmentRule object, and an error if there is any.
func (c *FakeClusterPodAssignmentRules) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterPodAssignmentRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterpodassignmentrulesResource, name), &v1alpha1.ClusterPodAssignmentRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPodAssignmentRule), err
}

// List takes label and field selectors, and returns the list of ClusterPodAssignmentRules that match those selectors.
func (c *FakeClusterPodAssignmentRules) List(opts v1.ListOptions) (result *v1alpha1.ClusterPodAssignmentRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterpodassignmentrulesResource, clusterpodassignmentrulesKind, opts), &v1alpha1.ClusterPodAssignmentRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterPodAssignmentRuleList{}
	for _, item := range obj.(*v1alpha1.ClusterPodAssignmentRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPodAssignmentRules.
func (c *FakeClusterPodAssignmentRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterpodassignmentrulesResource, opts))
}

// Create takes the representation of a clusterPodAssignmentRule and creates it.  Returns the server's representation of the clusterPodAssignmentRule, and an error, if there is any.
func (c *FakeClusterPodAssignmentRules) Create(clusterPodAssignmentRule *v1alpha1.ClusterPodAssignmentRule) (result *v1alpha1.ClusterPodAssignmentRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterpodassignmentrulesResource, clusterPodAssignmentRule), &v1alpha1.ClusterPodAssignmentRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPodAssignmentRule), err
}

// Update takes the representation of a clusterPodAssignmentRule and updates it. Returns the server's representation of the clusterPodAssignmentRule, and an error, if there is any.
func (c *FakeClusterPodAssignmentRules) Update(clusterPodAssignmentRule *v1alpha1.ClusterPodAssignmentRule) (result *v1alpha1.ClusterPodAssignmentRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterpodassignmentrulesResource, clusterPodAssignmentRule), &v1alpha1.ClusterPodAssignmentRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPodAssignmentRule), err
}

// Delete takes name of the clusterPodAssignmentRule and deletes it. Returns an error if one occurs.
func (c *FakeClusterPodAssignmentRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterpodassignmentrulesResource, name), &v1alpha1.ClusterPodAssignmentRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPodAssignmentRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterpodassignmentrulesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterPodAssignmentRuleList{})
	return err
}

// Patch applies the patch and returns the patched clusterPodAssignmentRule.
func (c *FakeClusterPodAssignmentRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterPodAssignmentRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterpodassignmentrulesResource, name, data, subresources...), &v1alpha1.ClusterPodAssignmentRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPodAssignmentRule), err
}
