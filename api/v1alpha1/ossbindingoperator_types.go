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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OssBindingOperatorSpec defines the desired state of OssBindingOperator
type OssBindingOperatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of OssBindingOperator. Edit ossbindingoperator_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// OssBindingOperatorStatus defines the observed state of OssBindingOperator
type OssBindingOperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// OssBindingOperator is the Schema for the ossbindingoperators API
type OssBindingOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OssBindingOperatorSpec   `json:"spec,omitempty"`
	Status OssBindingOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OssBindingOperatorList contains a list of OssBindingOperator
type OssBindingOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OssBindingOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OssBindingOperator{}, &OssBindingOperatorList{})
}
