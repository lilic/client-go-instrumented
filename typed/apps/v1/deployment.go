/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"github.com/lilic/client-go-instrumented/metrics"
	v1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	typedv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// DeploymentsGetter has a method to return a DeploymentInterface.
// A group's client should implement this interface.
type DeploymentsGetter interface {
	Deployments(namespace string) DeploymentInterface
}

// DeploymentInterface has methods to work with Deployment resources.
type DeploymentInterface interface {
	Create(*v1.Deployment) (*v1.Deployment, error)
	Update(*v1.Deployment) (*v1.Deployment, error)
	UpdateStatus(*v1.Deployment) (*v1.Deployment, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Deployment, error)
	List(opts metav1.ListOptions) (*v1.DeploymentList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Deployment, err error)
	GetScale(deploymentName string, options metav1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(deploymentName string, scale *autoscalingv1.Scale) (*autoscalingv1.Scale, error)

	DeploymentExpansion
}

// deployments implements DeploymentInterface
type deployments struct {
	deploymentInterface typedv1.DeploymentInterface
	clientMetrics       *metrics.ClientMetrics
	ns                  string
}

// newDeployments returns a Deployments
func newDeployments(namespace string, d typedv1.DeploymentInterface, m *metrics.ClientMetrics) *deployments {
	return &deployments{
		deploymentInterface: d,
		clientMetrics:       m,
		ns:                  namespace,
	}
}

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *deployments) Get(name string, options metav1.GetOptions) (result *v1.Deployment, err error) {
	result = &v1.Deployment{}
	result, err = c.deploymentInterface.Get(name, options)
	if c.clientMetrics != nil {
		c.clientMetrics.Inc("deployments", "get", name, c.ns, err)
	}
	return
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *deployments) List(opts metav1.ListOptions) (result *v1.DeploymentList, err error) {
	result = &v1.DeploymentList{}
	result, err = c.deploymentInterface.List(opts)
	if c.clientMetrics != nil {
		c.clientMetrics.Inc("deployments", "list", name, c.ns, err)
	}
	return
}

// Watch returns a watch.Interface that watches the requested deployments.
func (c *deployments) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	watch, err := c.deploymentInterface.Watch(opts)
	if c.clientMetrics != nil {
		c.clientMetrics.Inc("deployments", "watch", name, c.ns, err)
	}
	if err != nil {
		return nil, err
	}
	return watch, err
}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *deployments) Create(deployment *v1.Deployment) (result *v1.Deployment, err error) {
	result = &v1.Deployment{}
	result, err = c.deploymentInterface.Create(deployment)
	if c.clientMetrics != nil {
		c.clientMetrics.Inc("deployments", "create", name, c.ns, err)
	}
	return
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *deployments) Update(deployment *v1.Deployment) (result *v1.Deployment, err error) {
	result = &v1.Deployment{}
	result, err = c.deploymentInterface.Update(deployment)
	if c.clientMetrics != nil {
		c.clientMetrics.Inc("deployments", "update", name, c.ns, err)
	}
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *deployments) UpdateStatus(deployment *v1.Deployment) (result *v1.Deployment, err error) {
	result = &v1.Deployment{}
	result, err = c.deploymentInterface.UpdateStatus(deployment)
	if c.clientMetrics != nil {
		c.clientMetrics.Inc("deployments", "updatestatus", name, c.ns, err)
	}
	return
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *deployments) Delete(name string, options *metav1.DeleteOptions) error {
	err := c.deploymentInterface.Delete(name, options)
	if c.clientMetrics != nil {
		c.clientMetrics.Inc("deployments", "delete", name, c.ns, err)
	}
	if err != nil {
		return err
	}
	return nil
}

// DeleteCollection deletes a collection of objects.
func (c *deployments) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	err := c.deploymentInterface.DeleteCollection(options, listOptions)
	if c.clientMetrics != nil {
		c.clientMetrics.Inc("deployments", "deletecollection", name, c.ns, err)
	}
	if err != nil {
		return err
	}
	return nil
}

// Patch applies the patch and returns the patched deployment.
func (c *deployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Deployment, err error) {
	result = &v1.Deployment{}
	result, err = c.deploymentInterface.Patch(name, pt, data, subresources...)
	if c.clientMetrics != nil {
		c.clientMetrics.Inc("deployments", "patch", name, c.ns, err)
	}
	return
}

// GetScale takes name of the deployment, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *deployments) GetScale(deploymentName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	result, err = c.deploymentInterface.GetScale(deploymentName, options)
	if c.clientMetrics != nil {
		c.clientMetrics.Inc("deployments", "getscale", name, c.ns, err)
	}
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *deployments) UpdateScale(deploymentName string, scale *autoscalingv1.Scale) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	result, err = c.deploymentInterface.UpdateScale(deploymentName, options)
	if c.clientMetrics != nil {
		c.clientMetrics.Inc("deployments", "updatescale", name, c.ns, err)
	}
	return
}