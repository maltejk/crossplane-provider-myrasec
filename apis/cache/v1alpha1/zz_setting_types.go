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

type SettingObservation struct {
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Modified *string `json:"modified,omitempty" tf:"modified,omitempty"`

	SettingID *int64 `json:"settingId,omitempty" tf:"setting_id,omitempty"`
}

type SettingParameters struct {

	// Define wether this cache setting is enabled or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Enforce cache TTL allows you to set the cache TTL (Cache Control: max-age) in the backend regardless of the response sent from your Origin.
	// +kubebuilder:validation:Optional
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// How long an object will be cached. Origin responses with the HTTP codes 404 will be cached.
	// +kubebuilder:validation:Required
	NotFoundTTL *int64 `json:"notFoundTtl" tf:"not_found_ttl,omitempty"`

	// Path which must match to cache response.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// The order in which the cache rules take action as long as the cache sorting is activated.
	// +kubebuilder:validation:Optional
	Sort *int64 `json:"sort,omitempty" tf:"sort,omitempty"`

	// The Subdomain for the cache setting.
	// +kubebuilder:validation:Required
	SubdomainName *string `json:"subdomainName" tf:"subdomain_name,omitempty"`

	// Time to live.
	// +kubebuilder:validation:Required
	TTL *int64 `json:"ttl" tf:"ttl,omitempty"`

	// Type how path should match.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// SettingSpec defines the desired state of Setting
type SettingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SettingParameters `json:"forProvider"`
}

// SettingStatus defines the observed state of Setting.
type SettingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SettingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Setting is the Schema for the Settings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,myrasecjet}
type Setting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SettingSpec   `json:"spec"`
	Status            SettingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SettingList contains a list of Settings
type SettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Setting `json:"items"`
}

// Repository type metadata.
var (
	Setting_Kind             = "Setting"
	Setting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Setting_Kind}.String()
	Setting_KindAPIVersion   = Setting_Kind + "." + CRDGroupVersion.String()
	Setting_GroupVersionKind = CRDGroupVersion.WithKind(Setting_Kind)
)

func init() {
	SchemeBuilder.Register(&Setting{}, &SettingList{})
}