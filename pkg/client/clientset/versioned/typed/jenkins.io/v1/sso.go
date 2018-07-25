// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jenkins-x/sso-operator/pkg/apis/jenkins.io/v1"
	scheme "github.com/jenkins-x/sso-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SSOsGetter has a method to return a SSOInterface.
// A group's client should implement this interface.
type SSOsGetter interface {
	SSOs(namespace string) SSOInterface
}

// SSOInterface has methods to work with SSO resources.
type SSOInterface interface {
	Create(*v1.SSO) (*v1.SSO, error)
	Update(*v1.SSO) (*v1.SSO, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.SSO, error)
	List(opts metav1.ListOptions) (*v1.SSOList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SSO, err error)
	SSOExpansion
}

// sSOs implements SSOInterface
type sSOs struct {
	client rest.Interface
	ns     string
}

// newSSOs returns a SSOs
func newSSOs(c *JenkinsV1Client, namespace string) *sSOs {
	return &sSOs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sSO, and returns the corresponding sSO object, and an error if there is any.
func (c *sSOs) Get(name string, options metav1.GetOptions) (result *v1.SSO, err error) {
	result = &v1.SSO{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ssos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SSOs that match those selectors.
func (c *sSOs) List(opts metav1.ListOptions) (result *v1.SSOList, err error) {
	result = &v1.SSOList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ssos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sSOs.
func (c *sSOs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ssos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a sSO and creates it.  Returns the server's representation of the sSO, and an error, if there is any.
func (c *sSOs) Create(sSO *v1.SSO) (result *v1.SSO, err error) {
	result = &v1.SSO{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ssos").
		Body(sSO).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sSO and updates it. Returns the server's representation of the sSO, and an error, if there is any.
func (c *sSOs) Update(sSO *v1.SSO) (result *v1.SSO, err error) {
	result = &v1.SSO{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ssos").
		Name(sSO.Name).
		Body(sSO).
		Do().
		Into(result)
	return
}

// Delete takes name of the sSO and deletes it. Returns an error if one occurs.
func (c *sSOs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ssos").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sSOs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ssos").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sSO.
func (c *sSOs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SSO, err error) {
	result = &v1.SSO{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ssos").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}