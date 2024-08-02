/*
Copyright 2024.

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

package v1alpha1

import (
	v1alpha1 "inftyai.com/llmaz/api/inference/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// PlaygroundLister helps list Playgrounds.
// All objects returned here must be treated as read-only.
type PlaygroundLister interface {
	// List lists all Playgrounds in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Playground, err error)
	// Playgrounds returns an object that can list and get Playgrounds.
	Playgrounds(namespace string) PlaygroundNamespaceLister
	PlaygroundListerExpansion
}

// playgroundLister implements the PlaygroundLister interface.
type playgroundLister struct {
	listers.ResourceIndexer[*v1alpha1.Playground]
}

// NewPlaygroundLister returns a new PlaygroundLister.
func NewPlaygroundLister(indexer cache.Indexer) PlaygroundLister {
	return &playgroundLister{listers.New[*v1alpha1.Playground](indexer, v1alpha1.Resource("playground"))}
}

// Playgrounds returns an object that can list and get Playgrounds.
func (s *playgroundLister) Playgrounds(namespace string) PlaygroundNamespaceLister {
	return playgroundNamespaceLister{listers.NewNamespaced[*v1alpha1.Playground](s.ResourceIndexer, namespace)}
}

// PlaygroundNamespaceLister helps list and get Playgrounds.
// All objects returned here must be treated as read-only.
type PlaygroundNamespaceLister interface {
	// List lists all Playgrounds in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Playground, err error)
	// Get retrieves the Playground from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Playground, error)
	PlaygroundNamespaceListerExpansion
}

// playgroundNamespaceLister implements the PlaygroundNamespaceLister
// interface.
type playgroundNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.Playground]
}
