/*
Copyright 2022.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/ihcsim/cbt-aggapi/pkg/apis/discovery/v1alpha1"
	scheme "github.com/ihcsim/cbt-aggapi/pkg/generated/discovery/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CSIDriverEndpointsGetter has a method to return a CSIDriverEndpointInterface.
// A group's client should implement this interface.
type CSIDriverEndpointsGetter interface {
	CSIDriverEndpoints(namespace string) CSIDriverEndpointInterface
}

// CSIDriverEndpointInterface has methods to work with CSIDriverEndpoint resources.
type CSIDriverEndpointInterface interface {
	Create(ctx context.Context, cSIDriverEndpoint *v1alpha1.CSIDriverEndpoint, opts v1.CreateOptions) (*v1alpha1.CSIDriverEndpoint, error)
	Update(ctx context.Context, cSIDriverEndpoint *v1alpha1.CSIDriverEndpoint, opts v1.UpdateOptions) (*v1alpha1.CSIDriverEndpoint, error)
	UpdateStatus(ctx context.Context, cSIDriverEndpoint *v1alpha1.CSIDriverEndpoint, opts v1.UpdateOptions) (*v1alpha1.CSIDriverEndpoint, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CSIDriverEndpoint, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CSIDriverEndpointList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CSIDriverEndpoint, err error)
	CSIDriverEndpointExpansion
}

// cSIDriverEndpoints implements CSIDriverEndpointInterface
type cSIDriverEndpoints struct {
	client rest.Interface
	ns     string
}

// newCSIDriverEndpoints returns a CSIDriverEndpoints
func newCSIDriverEndpoints(c *DiscoveryV1alpha1Client, namespace string) *cSIDriverEndpoints {
	return &cSIDriverEndpoints{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cSIDriverEndpoint, and returns the corresponding cSIDriverEndpoint object, and an error if there is any.
func (c *cSIDriverEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CSIDriverEndpoint, err error) {
	result = &v1alpha1.CSIDriverEndpoint{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("csidriverendpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CSIDriverEndpoints that match those selectors.
func (c *cSIDriverEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CSIDriverEndpointList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CSIDriverEndpointList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("csidriverendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cSIDriverEndpoints.
func (c *cSIDriverEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("csidriverendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cSIDriverEndpoint and creates it.  Returns the server's representation of the cSIDriverEndpoint, and an error, if there is any.
func (c *cSIDriverEndpoints) Create(ctx context.Context, cSIDriverEndpoint *v1alpha1.CSIDriverEndpoint, opts v1.CreateOptions) (result *v1alpha1.CSIDriverEndpoint, err error) {
	result = &v1alpha1.CSIDriverEndpoint{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("csidriverendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cSIDriverEndpoint).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cSIDriverEndpoint and updates it. Returns the server's representation of the cSIDriverEndpoint, and an error, if there is any.
func (c *cSIDriverEndpoints) Update(ctx context.Context, cSIDriverEndpoint *v1alpha1.CSIDriverEndpoint, opts v1.UpdateOptions) (result *v1alpha1.CSIDriverEndpoint, err error) {
	result = &v1alpha1.CSIDriverEndpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("csidriverendpoints").
		Name(cSIDriverEndpoint.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cSIDriverEndpoint).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cSIDriverEndpoints) UpdateStatus(ctx context.Context, cSIDriverEndpoint *v1alpha1.CSIDriverEndpoint, opts v1.UpdateOptions) (result *v1alpha1.CSIDriverEndpoint, err error) {
	result = &v1alpha1.CSIDriverEndpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("csidriverendpoints").
		Name(cSIDriverEndpoint.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cSIDriverEndpoint).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cSIDriverEndpoint and deletes it. Returns an error if one occurs.
func (c *cSIDriverEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("csidriverendpoints").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cSIDriverEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("csidriverendpoints").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cSIDriverEndpoint.
func (c *cSIDriverEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CSIDriverEndpoint, err error) {
	result = &v1alpha1.CSIDriverEndpoint{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("csidriverendpoints").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
