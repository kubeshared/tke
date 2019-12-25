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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	auth "tkestack.io/tke/api/auth"
)

// LocalIdentityLister helps list LocalIdentities.
type LocalIdentityLister interface {
	// List lists all LocalIdentities in the indexer.
	List(selector labels.Selector) (ret []*auth.LocalIdentity, err error)
	// Get retrieves the LocalIdentity from the index for a given name.
	Get(name string) (*auth.LocalIdentity, error)
	LocalIdentityListerExpansion
}

// localIdentityLister implements the LocalIdentityLister interface.
type localIdentityLister struct {
	indexer cache.Indexer
}

// NewLocalIdentityLister returns a new LocalIdentityLister.
func NewLocalIdentityLister(indexer cache.Indexer) LocalIdentityLister {
	return &localIdentityLister{indexer: indexer}
}

// List lists all LocalIdentities in the indexer.
func (s *localIdentityLister) List(selector labels.Selector) (ret []*auth.LocalIdentity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*auth.LocalIdentity))
	})
	return ret, err
}

// Get retrieves the LocalIdentity from the index for a given name.
func (s *localIdentityLister) Get(name string) (*auth.LocalIdentity, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(auth.Resource("localidentity"), name)
	}
	return obj.(*auth.LocalIdentity), nil
}
