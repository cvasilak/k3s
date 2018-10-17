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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	auditregistration "k8s.io/kubernetes/pkg/apis/auditregistration"
)

// AuditSinkLister helps list AuditSinks.
type AuditSinkLister interface {
	// List lists all AuditSinks in the indexer.
	List(selector labels.Selector) (ret []*auditregistration.AuditSink, err error)
	// Get retrieves the AuditSink from the index for a given name.
	Get(name string) (*auditregistration.AuditSink, error)
	AuditSinkListerExpansion
}

// auditSinkLister implements the AuditSinkLister interface.
type auditSinkLister struct {
	indexer cache.Indexer
}

// NewAuditSinkLister returns a new AuditSinkLister.
func NewAuditSinkLister(indexer cache.Indexer) AuditSinkLister {
	return &auditSinkLister{indexer: indexer}
}

// List lists all AuditSinks in the indexer.
func (s *auditSinkLister) List(selector labels.Selector) (ret []*auditregistration.AuditSink, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*auditregistration.AuditSink))
	})
	return ret, err
}

// Get retrieves the AuditSink from the index for a given name.
func (s *auditSinkLister) Get(name string) (*auditregistration.AuditSink, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(auditregistration.Resource("auditsink"), name)
	}
	return obj.(*auditregistration.AuditSink), nil
}
