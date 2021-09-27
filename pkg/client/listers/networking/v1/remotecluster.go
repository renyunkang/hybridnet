/*
Copyright 2021 The Hybridnet Authors.

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

package v1

import (
	v1 "github.com/alibaba/hybridnet/pkg/apis/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RemoteClusterLister helps list RemoteClusters.
type RemoteClusterLister interface {
	// List lists all RemoteClusters in the indexer.
	List(selector labels.Selector) (ret []*v1.RemoteCluster, err error)
	// Get retrieves the RemoteCluster from the index for a given name.
	Get(name string) (*v1.RemoteCluster, error)
	RemoteClusterListerExpansion
}

// remoteClusterLister implements the RemoteClusterLister interface.
type remoteClusterLister struct {
	indexer cache.Indexer
}

// NewRemoteClusterLister returns a new RemoteClusterLister.
func NewRemoteClusterLister(indexer cache.Indexer) RemoteClusterLister {
	return &remoteClusterLister{indexer: indexer}
}

// List lists all RemoteClusters in the indexer.
func (s *remoteClusterLister) List(selector labels.Selector) (ret []*v1.RemoteCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.RemoteCluster))
	})
	return ret, err
}

// Get retrieves the RemoteCluster from the index for a given name.
func (s *remoteClusterLister) Get(name string) (*v1.RemoteCluster, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("remotecluster"), name)
	}
	return obj.(*v1.RemoteCluster), nil
}
