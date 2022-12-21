/*
Copyright 2022 Akash Pathak <pathakvikash9211@gmail.com>.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CostForecasterSpec defines the desired state of CostForecaster
type CostForecasterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CostForecaster. Edit costforecaster_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// CostForecasterStatus defines the observed state of CostForecaster
type CostForecasterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CostForecaster is the Schema for the costforecasters API
type CostForecaster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CostForecasterSpec   `json:"spec,omitempty"`
	Status CostForecasterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CostForecasterList contains a list of CostForecaster
type CostForecasterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CostForecaster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CostForecaster{}, &CostForecasterList{})
}
