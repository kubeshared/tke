/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	auth "tkestack.io/tke/api/auth"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
)

// IdentityProvidersGetter has a method to return a IdentityProviderInterface.
// A group's client should implement this interface.
type IdentityProvidersGetter interface {
	IdentityProviders() IdentityProviderInterface
}

// IdentityProviderInterface has methods to work with IdentityProvider resources.
type IdentityProviderInterface interface {
	Create(*auth.IdentityProvider) (*auth.IdentityProvider, error)
	Update(*auth.IdentityProvider) (*auth.IdentityProvider, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*auth.IdentityProvider, error)
	List(opts v1.ListOptions) (*auth.IdentityProviderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auth.IdentityProvider, err error)
	IdentityProviderExpansion
}

// identityProviders implements IdentityProviderInterface
type identityProviders struct {
	client rest.Interface
}

// newIdentityProviders returns a IdentityProviders
func newIdentityProviders(c *AuthClient) *identityProviders {
	return &identityProviders{
		client: c.RESTClient(),
	}
}

// Get takes name of the identityProvider, and returns the corresponding identityProvider object, and an error if there is any.
func (c *identityProviders) Get(name string, options v1.GetOptions) (result *auth.IdentityProvider, err error) {
	result = &auth.IdentityProvider{}
	err = c.client.Get().
		Resource("identityproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IdentityProviders that match those selectors.
func (c *identityProviders) List(opts v1.ListOptions) (result *auth.IdentityProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &auth.IdentityProviderList{}
	err = c.client.Get().
		Resource("identityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested identityProviders.
func (c *identityProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("identityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a identityProvider and creates it.  Returns the server's representation of the identityProvider, and an error, if there is any.
func (c *identityProviders) Create(identityProvider *auth.IdentityProvider) (result *auth.IdentityProvider, err error) {
	result = &auth.IdentityProvider{}
	err = c.client.Post().
		Resource("identityproviders").
		Body(identityProvider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a identityProvider and updates it. Returns the server's representation of the identityProvider, and an error, if there is any.
func (c *identityProviders) Update(identityProvider *auth.IdentityProvider) (result *auth.IdentityProvider, err error) {
	result = &auth.IdentityProvider{}
	err = c.client.Put().
		Resource("identityproviders").
		Name(identityProvider.Name).
		Body(identityProvider).
		Do().
		Into(result)
	return
}

// Delete takes name of the identityProvider and deletes it. Returns an error if one occurs.
func (c *identityProviders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("identityproviders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *identityProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("identityproviders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched identityProvider.
func (c *identityProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auth.IdentityProvider, err error) {
	result = &auth.IdentityProvider{}
	err = c.client.Patch(pt).
		Resource("identityproviders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
