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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// BookSpecApplyConfiguration represents an declarative configuration of the BookSpec type for use
// with apply.
type BookSpecApplyConfiguration struct {
	Name      *string                          `json:"name,omitempty"`
	Replicas  *int32                           `json:"replicas,omitempty"`
	Container *ContainerSpecApplyConfiguration `json:"container,omitempty"`
}

// BookSpecApplyConfiguration constructs an declarative configuration of the BookSpec type for use with
// apply.
func BookSpec() *BookSpecApplyConfiguration {
	return &BookSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *BookSpecApplyConfiguration) WithName(value string) *BookSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *BookSpecApplyConfiguration) WithReplicas(value int32) *BookSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithContainer sets the Container field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Container field is set to the value of the last call.
func (b *BookSpecApplyConfiguration) WithContainer(value *ContainerSpecApplyConfiguration) *BookSpecApplyConfiguration {
	b.Container = value
	return b
}
