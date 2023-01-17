// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/openshift/api/config/v1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OperatorHubsGetter has a method to return a OperatorHubInterface.
// A group's client should implement this interface.
type OperatorHubsGetter interface {
	OperatorHubs() OperatorHubInterface
}

// OperatorHubInterface has methods to work with OperatorHub resources.
type OperatorHubInterface interface {
	Create(ctx context.Context, operatorHub *v1.OperatorHub, opts metav1.CreateOptions) (*v1.OperatorHub, error)
	Update(ctx context.Context, operatorHub *v1.OperatorHub, opts metav1.UpdateOptions) (*v1.OperatorHub, error)
	UpdateStatus(ctx context.Context, operatorHub *v1.OperatorHub, opts metav1.UpdateOptions) (*v1.OperatorHub, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.OperatorHub, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.OperatorHubList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OperatorHub, err error)
	Apply(ctx context.Context, operatorHub *configv1.OperatorHubApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OperatorHub, err error)
	ApplyStatus(ctx context.Context, operatorHub *configv1.OperatorHubApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OperatorHub, err error)
	OperatorHubExpansion
}

// operatorHubs implements OperatorHubInterface
type operatorHubs struct {
	client rest.Interface
}

// newOperatorHubs returns a OperatorHubs
func newOperatorHubs(c *ConfigV1Client) *operatorHubs {
	return &operatorHubs{
		client: c.RESTClient(),
	}
}

// Get takes name of the operatorHub, and returns the corresponding operatorHub object, and an error if there is any.
func (c *operatorHubs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.OperatorHub, err error) {
	result = &v1.OperatorHub{}
	err = c.client.Get().
		Resource("operatorhubs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OperatorHubs that match those selectors.
func (c *operatorHubs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OperatorHubList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.OperatorHubList{}
	err = c.client.Get().
		Resource("operatorhubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested operatorHubs.
func (c *operatorHubs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("operatorhubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a operatorHub and creates it.  Returns the server's representation of the operatorHub, and an error, if there is any.
func (c *operatorHubs) Create(ctx context.Context, operatorHub *v1.OperatorHub, opts metav1.CreateOptions) (result *v1.OperatorHub, err error) {
	result = &v1.OperatorHub{}
	err = c.client.Post().
		Resource("operatorhubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operatorHub).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a operatorHub and updates it. Returns the server's representation of the operatorHub, and an error, if there is any.
func (c *operatorHubs) Update(ctx context.Context, operatorHub *v1.OperatorHub, opts metav1.UpdateOptions) (result *v1.OperatorHub, err error) {
	result = &v1.OperatorHub{}
	err = c.client.Put().
		Resource("operatorhubs").
		Name(operatorHub.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operatorHub).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *operatorHubs) UpdateStatus(ctx context.Context, operatorHub *v1.OperatorHub, opts metav1.UpdateOptions) (result *v1.OperatorHub, err error) {
	result = &v1.OperatorHub{}
	err = c.client.Put().
		Resource("operatorhubs").
		Name(operatorHub.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operatorHub).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the operatorHub and deletes it. Returns an error if one occurs.
func (c *operatorHubs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("operatorhubs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *operatorHubs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("operatorhubs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched operatorHub.
func (c *operatorHubs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OperatorHub, err error) {
	result = &v1.OperatorHub{}
	err = c.client.Patch(pt).
		Resource("operatorhubs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied operatorHub.
func (c *operatorHubs) Apply(ctx context.Context, operatorHub *configv1.OperatorHubApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OperatorHub, err error) {
	if operatorHub == nil {
		return nil, fmt.Errorf("operatorHub provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(operatorHub)
	if err != nil {
		return nil, err
	}
	name := operatorHub.Name
	if name == nil {
		return nil, fmt.Errorf("operatorHub.Name must be provided to Apply")
	}
	result = &v1.OperatorHub{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("operatorhubs").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *operatorHubs) ApplyStatus(ctx context.Context, operatorHub *configv1.OperatorHubApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OperatorHub, err error) {
	if operatorHub == nil {
		return nil, fmt.Errorf("operatorHub provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(operatorHub)
	if err != nil {
		return nil, err
	}

	name := operatorHub.Name
	if name == nil {
		return nil, fmt.Errorf("operatorHub.Name must be provided to Apply")
	}

	result = &v1.OperatorHub{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("operatorhubs").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
