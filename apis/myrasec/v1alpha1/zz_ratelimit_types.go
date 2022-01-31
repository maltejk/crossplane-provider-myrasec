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

type RatelimitObservation struct {
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Modified *string `json:"modified,omitempty" tf:"modified,omitempty"`

	RatelimitID *int64 `json:"ratelimitId,omitempty" tf:"ratelimit_id,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RatelimitParameters struct {

	// Burst defines how many requests a client can make in excess of the specified rate.
	// +kubebuilder:validation:Optional
	Burst *int64 `json:"burst,omitempty" tf:"burst,omitempty"`

	// Network in CIDR notation affected by the rate limiter.
	// +kubebuilder:validation:Required
	Network *string `json:"network" tf:"network,omitempty"`

	// The Subdomain for the rate limit setting.
	// +kubebuilder:validation:Required
	SubdomainName *string `json:"subdomainName" tf:"subdomain_name,omitempty"`

	// The affected timeframe in seconds for the rate limit.
	// +kubebuilder:validation:Optional
	Timeframe *int64 `json:"timeframe,omitempty" tf:"timeframe,omitempty"`

	// Maximum amount of requests for the given network.
	// +kubebuilder:validation:Optional
	Value *int64 `json:"value,omitempty" tf:"value,omitempty"`
}

// RatelimitSpec defines the desired state of Ratelimit
type RatelimitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RatelimitParameters `json:"forProvider"`
}

// RatelimitStatus defines the observed state of Ratelimit.
type RatelimitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RatelimitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ratelimit is the Schema for the Ratelimits API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,myrasecjet}
type Ratelimit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RatelimitSpec   `json:"spec"`
	Status            RatelimitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RatelimitList contains a list of Ratelimits
type RatelimitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ratelimit `json:"items"`
}

// Repository type metadata.
var (
	Ratelimit_Kind             = "Ratelimit"
	Ratelimit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Ratelimit_Kind}.String()
	Ratelimit_KindAPIVersion   = Ratelimit_Kind + "." + CRDGroupVersion.String()
	Ratelimit_GroupVersionKind = CRDGroupVersion.WithKind(Ratelimit_Kind)
)

func init() {
	SchemeBuilder.Register(&Ratelimit{}, &RatelimitList{})
}
