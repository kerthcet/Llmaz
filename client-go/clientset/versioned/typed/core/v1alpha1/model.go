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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "inftyai.com/llmaz/api/core/v1alpha1"
	corev1alpha1 "inftyai.com/llmaz/client-go/applyconfiguration/core/v1alpha1"
	scheme "inftyai.com/llmaz/client-go/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ModelsGetter has a method to return a ModelInterface.
// A group's client should implement this interface.
type ModelsGetter interface {
	Models(namespace string) ModelInterface
}

// ModelInterface has methods to work with Model resources.
type ModelInterface interface {
	Create(ctx context.Context, model *v1alpha1.Model, opts v1.CreateOptions) (*v1alpha1.Model, error)
	Update(ctx context.Context, model *v1alpha1.Model, opts v1.UpdateOptions) (*v1alpha1.Model, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, model *v1alpha1.Model, opts v1.UpdateOptions) (*v1alpha1.Model, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Model, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ModelList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Model, err error)
	Apply(ctx context.Context, model *corev1alpha1.ModelApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Model, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, model *corev1alpha1.ModelApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Model, err error)
	ModelExpansion
}

// models implements ModelInterface
type models struct {
	*gentype.ClientWithListAndApply[*v1alpha1.Model, *v1alpha1.ModelList, *corev1alpha1.ModelApplyConfiguration]
}

// newModels returns a Models
func newModels(c *LlmazV1alpha1Client, namespace string) *models {
	return &models{
		gentype.NewClientWithListAndApply[*v1alpha1.Model, *v1alpha1.ModelList, *corev1alpha1.ModelApplyConfiguration](
			"models",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.Model { return &v1alpha1.Model{} },
			func() *v1alpha1.ModelList { return &v1alpha1.ModelList{} }),
	}
}
