/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CoffeeObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Price *int64 `json:"price,omitempty" tf:"price,omitempty"`

	Teaser *string `json:"teaser,omitempty" tf:"teaser,omitempty"`
}

type CoffeeParameters struct {

	// +kubebuilder:validation:Required
	ID *int64 `json:"id" tf:"id,omitempty"`
}

type ItemsObservation struct {
}

type ItemsParameters struct {

	// +kubebuilder:validation:Required
	Coffee []CoffeeParameters `json:"coffee" tf:"coffee,omitempty"`

	// +kubebuilder:validation:Required
	Quantity *int64 `json:"quantity" tf:"quantity,omitempty"`
}

type OrderObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OrderParameters struct {

	// +kubebuilder:validation:Required
	Items []ItemsParameters `json:"items" tf:"items,omitempty"`

	// +kubebuilder:validation:Optional
	LastUpdated *string `json:"lastUpdated,omitempty" tf:"last_updated,omitempty"`
}

// OrderSpec defines the desired state of Order
type OrderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrderParameters `json:"forProvider"`
}

// OrderStatus defines the observed state of Order.
type OrderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Order is the Schema for the Orders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,templatejet}
type Order struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrderSpec   `json:"spec"`
	Status            OrderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrderList contains a list of Orders
type OrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Order `json:"items"`
}

// Repository type metadata.
var (
	Order_Kind             = "Order"
	Order_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Order_Kind}.String()
	Order_KindAPIVersion   = Order_Kind + "." + CRDGroupVersion.String()
	Order_GroupVersionKind = CRDGroupVersion.WithKind(Order_Kind)
)

func init() {
	SchemeBuilder.Register(&Order{}, &OrderList{})
}
