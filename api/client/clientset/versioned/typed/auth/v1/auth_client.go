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

package v1

import (
	rest "k8s.io/client-go/rest"
	v1 "tkestack.io/tke/api/auth/v1"
	"tkestack.io/tke/api/client/clientset/versioned/scheme"
)

type AuthV1Interface interface {
	RESTClient() rest.Interface
	APIKeysGetter
	APISigningKeysGetter
	CategoriesGetter
	ClientsGetter
	ConfigMapsGetter
	GroupsGetter
	IdentityProvidersGetter
	LocalGroupsGetter
	LocalIdentitiesGetter
	PoliciesGetter
	RolesGetter
	RulesGetter
	UsersGetter
}

// AuthV1Client is used to interact with features provided by the auth.tkestack.io group.
type AuthV1Client struct {
	restClient rest.Interface
}

func (c *AuthV1Client) APIKeys() APIKeyInterface {
	return newAPIKeys(c)
}

func (c *AuthV1Client) APISigningKeys() APISigningKeyInterface {
	return newAPISigningKeys(c)
}

func (c *AuthV1Client) Categories() CategoryInterface {
	return newCategories(c)
}

func (c *AuthV1Client) Clients() ClientInterface {
	return newClients(c)
}

func (c *AuthV1Client) ConfigMaps() ConfigMapInterface {
	return newConfigMaps(c)
}

func (c *AuthV1Client) Groups() GroupInterface {
	return newGroups(c)
}

func (c *AuthV1Client) IdentityProviders() IdentityProviderInterface {
	return newIdentityProviders(c)
}

func (c *AuthV1Client) LocalGroups() LocalGroupInterface {
	return newLocalGroups(c)
}

func (c *AuthV1Client) LocalIdentities() LocalIdentityInterface {
	return newLocalIdentities(c)
}

func (c *AuthV1Client) Policies() PolicyInterface {
	return newPolicies(c)
}

func (c *AuthV1Client) Roles() RoleInterface {
	return newRoles(c)
}

func (c *AuthV1Client) Rules() RuleInterface {
	return newRules(c)
}

func (c *AuthV1Client) Users() UserInterface {
	return newUsers(c)
}

// NewForConfig creates a new AuthV1Client for the given config.
func NewForConfig(c *rest.Config) (*AuthV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &AuthV1Client{client}, nil
}

// NewForConfigOrDie creates a new AuthV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AuthV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AuthV1Client for the given RESTClient.
func New(c rest.Interface) *AuthV1Client {
	return &AuthV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *AuthV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
