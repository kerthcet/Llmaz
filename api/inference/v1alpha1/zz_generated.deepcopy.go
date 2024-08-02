//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
corev1alpha1 "inftyai.com/llmaz/api/core/v1alpha1"
runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfig) DeepCopyInto(out *BackendConfig) {
*out = *in
if in.Name != nil {
in, out := &in.Name, &out.Name
*out = new(BackendName)
**out = **in
}
if in.Version != nil {
in, out := &in.Version, &out.Version
*out = new(string)
**out = **in
}
if in.Args != nil {
in, out := &in.Args, &out.Args
*out = make([]string, len(*in))
copy(*out, *in)
}
if in.Envs != nil {
in, out := &in.Envs, &out.Envs
*out = make([]invalid type, len(*in))
for i := range *in {
}
}
if in.Resources != nil {
in, out := &in.Resources, &out.Resources
*out = new(ResourceRequirements)
(*in).DeepCopyInto(*out)
}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfig.
func (in *BackendConfig) DeepCopy() *BackendConfig {
	if in == nil { return nil }
	out := new(BackendConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticConfig) DeepCopyInto(out *ElasticConfig) {
*out = *in
if in.MinReplicas != nil {
in, out := &in.MinReplicas, &out.MinReplicas
*out = new(int32)
**out = **in
}
if in.MaxReplicas != nil {
in, out := &in.MaxReplicas, &out.MaxReplicas
*out = new(int32)
**out = **in
}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticConfig.
func (in *ElasticConfig) DeepCopy() *ElasticConfig {
	if in == nil { return nil }
	out := new(ElasticConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Playground) DeepCopyInto(out *Playground) {
*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Playground.
func (in *Playground) DeepCopy() *Playground {
	if in == nil { return nil }
	out := new(Playground)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Playground) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaygroundList) DeepCopyInto(out *PlaygroundList) {
*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaygroundList.
func (in *PlaygroundList) DeepCopy() *PlaygroundList {
	if in == nil { return nil }
	out := new(PlaygroundList)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlaygroundList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaygroundSpec) DeepCopyInto(out *PlaygroundSpec) {
*out = *in
if in.Replicas != nil {
in, out := &in.Replicas, &out.Replicas
*out = new(int32)
**out = **in
}
if in.ModelClaim != nil {
in, out := &in.ModelClaim, &out.ModelClaim
*out = new(corev1alpha1.ModelClaim)
(*in).DeepCopyInto(*out)
}
if in.MultiModelsClaims != nil {
in, out := &in.MultiModelsClaims, &out.MultiModelsClaims
*out = make([]corev1alpha1.MultiModelsClaim, len(*in))
for i := range *in {
(*in)[i].DeepCopyInto(&(*out)[i])
}
}
if in.BackendConfig != nil {
in, out := &in.BackendConfig, &out.BackendConfig
*out = new(BackendConfig)
(*in).DeepCopyInto(*out)
}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaygroundSpec.
func (in *PlaygroundSpec) DeepCopy() *PlaygroundSpec {
	if in == nil { return nil }
	out := new(PlaygroundSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaygroundStatus) DeepCopyInto(out *PlaygroundStatus) {
*out = *in
if in.Conditions != nil {
in, out := &in.Conditions, &out.Conditions
*out = make([]invalid type, len(*in))
for i := range *in {
}
}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaygroundStatus.
func (in *PlaygroundStatus) DeepCopy() *PlaygroundStatus {
	if in == nil { return nil }
	out := new(PlaygroundStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirements.
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
	if in == nil { return nil }
	out := new(ResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil { return nil }
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Service) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceList) DeepCopyInto(out *ServiceList) {
*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceList.
func (in *ServiceList) DeepCopy() *ServiceList {
	if in == nil { return nil }
	out := new(ServiceList)
	in.DeepCopyInto(out)
	return out
}


// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
*out = *in
if in.MultiModelsClaims != nil {
in, out := &in.MultiModelsClaims, &out.MultiModelsClaims
*out = make([]corev1alpha1.MultiModelsClaim, len(*in))
for i := range *in {
(*in)[i].DeepCopyInto(&(*out)[i])
}
}
in.WorkloadTemplate.DeepCopyInto(&out.WorkloadTemplate)
if in.ElasticConfig != nil {
in, out := &in.ElasticConfig, &out.ElasticConfig
*out = new(ElasticConfig)
(*in).DeepCopyInto(*out)
}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil { return nil }
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
*out = *in
if in.Conditions != nil {
in, out := &in.Conditions, &out.Conditions
*out = make([]invalid type, len(*in))
for i := range *in {
}
}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceStatus.
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
	if in == nil { return nil }
	out := new(ServiceStatus)
	in.DeepCopyInto(out)
	return out
}

