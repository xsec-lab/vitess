/*
Copyright 2020 The Vitess Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/xsec-lab/go/vt/topo/k8stopo/apis/topo/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VitessTopoNodeLister helps list VitessTopoNodes.
type VitessTopoNodeLister interface {
	// List lists all VitessTopoNodes in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.VitessTopoNode, err error)
	// VitessTopoNodes returns an object that can list and get VitessTopoNodes.
	VitessTopoNodes(namespace string) VitessTopoNodeNamespaceLister
	VitessTopoNodeListerExpansion
}

// vitessTopoNodeLister implements the VitessTopoNodeLister interface.
type vitessTopoNodeLister struct {
	indexer cache.Indexer
}

// NewVitessTopoNodeLister returns a new VitessTopoNodeLister.
func NewVitessTopoNodeLister(indexer cache.Indexer) VitessTopoNodeLister {
	return &vitessTopoNodeLister{indexer: indexer}
}

// List lists all VitessTopoNodes in the indexer.
func (s *vitessTopoNodeLister) List(selector labels.Selector) (ret []*v1beta1.VitessTopoNode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VitessTopoNode))
	})
	return ret, err
}

// VitessTopoNodes returns an object that can list and get VitessTopoNodes.
func (s *vitessTopoNodeLister) VitessTopoNodes(namespace string) VitessTopoNodeNamespaceLister {
	return vitessTopoNodeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VitessTopoNodeNamespaceLister helps list and get VitessTopoNodes.
type VitessTopoNodeNamespaceLister interface {
	// List lists all VitessTopoNodes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.VitessTopoNode, err error)
	// Get retrieves the VitessTopoNode from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.VitessTopoNode, error)
	VitessTopoNodeNamespaceListerExpansion
}

// vitessTopoNodeNamespaceLister implements the VitessTopoNodeNamespaceLister
// interface.
type vitessTopoNodeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VitessTopoNodes in the indexer for a given namespace.
func (s vitessTopoNodeNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.VitessTopoNode, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VitessTopoNode))
	})
	return ret, err
}

// Get retrieves the VitessTopoNode from the indexer for a given namespace and name.
func (s vitessTopoNodeNamespaceLister) Get(name string) (*v1beta1.VitessTopoNode, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("vitesstoponode"), name)
	}
	return obj.(*v1beta1.VitessTopoNode), nil
}
