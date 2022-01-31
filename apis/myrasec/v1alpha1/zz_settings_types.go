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

type SettingsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SettingsParameters struct {

	// Activate separated access log
	// +kubebuilder:validation:Optional
	AccessLog *bool `json:"accessLog,omitempty" tf:"access_log,omitempty"`

	// Detection of POST floods by using a JavaScript based puzzle.
	// +kubebuilder:validation:Optional
	AntibotPostFlood *bool `json:"antibotPostFlood,omitempty" tf:"antibot_post_flood,omitempty"`

	// This parameter determines the frequency how often the puzzle has to be solved. The higher the value the less likely the puzzle needs to be solved.
	// +kubebuilder:validation:Optional
	AntibotPostFloodThreshold *int64 `json:"antibotPostFloodThreshold,omitempty" tf:"antibot_post_flood_threshold,omitempty"`

	// Detection of valid clients by using a JavaScript based puzzle.
	// +kubebuilder:validation:Optional
	AntibotProofOfWork *bool `json:"antibotProofOfWork,omitempty" tf:"antibot_proof_of_work,omitempty"`

	// This parameter determines the frequency how often the puzzle has to be solved. The higher the value the less likely the puzzle needs to be solved.
	// +kubebuilder:validation:Optional
	AntibotProofOfWorkThreshold *int64 `json:"antibotProofOfWorkThreshold,omitempty" tf:"antibot_proof_of_work_threshold,omitempty"`

	// Specifies with which method requests are balanced between upstream servers.
	// +kubebuilder:validation:Optional
	BalancingMethod *string `json:"balancingMethod,omitempty" tf:"balancing_method,omitempty"`

	// Block all IPs, which are not whitelisted.
	// +kubebuilder:validation:Optional
	BlockNotWhitelisted *bool `json:"blockNotWhitelisted,omitempty" tf:"block_not_whitelisted,omitempty"`

	// Block traffic from the TOR network.
	// +kubebuilder:validation:Optional
	BlockTorNetwork *bool `json:"blockTorNetwork,omitempty" tf:"block_tor_network,omitempty"`

	// Turn caching on or off.
	// +kubebuilder:validation:Optional
	CacheEnabled *bool `json:"cacheEnabled,omitempty" tf:"cache_enabled,omitempty"`

	// Enable stale cache item revalidation.
	// +kubebuilder:validation:Optional
	CacheRevalidate *bool `json:"cacheRevalidate,omitempty" tf:"cache_revalidate,omitempty"`

	// Use subdomain as Content Delivery Node (CDN).
	// +kubebuilder:validation:Optional
	Cdn *bool `json:"cdn,omitempty" tf:"cdn,omitempty"`

	// Sets the maximum allowed size of the client request body, specified in the “Content-Length” request header field. Maximum 100MB.
	// +kubebuilder:validation:Optional
	ClientMaxBodySize *int64 `json:"clientMaxBodySize,omitempty" tf:"client_max_body_size,omitempty"`

	// The Diffie-Hellman key exchange parameter length.
	// +kubebuilder:validation:Optional
	DiffieHellmanExchange *int64 `json:"diffieHellmanExchange,omitempty" tf:"diffie_hellman_exchange,omitempty"`

	// Enable or disable origin SNI.
	// +kubebuilder:validation:Optional
	EnableOriginSni *bool `json:"enableOriginSni,omitempty" tf:"enable_origin_sni,omitempty"`

	// Set your own X-Forwarded-For header.
	// +kubebuilder:validation:Optional
	ForwardedForReplacement *string `json:"forwardedForReplacement,omitempty" tf:"forwarded_for_replacement,omitempty"`

	// Allows to set a port for communication with origin via HTTP.
	// +kubebuilder:validation:Optional
	HTTPOriginPort *int64 `json:"httpOriginPort,omitempty" tf:"http_origin_port,omitempty"`

	// HSTS Strict Transport Security (HSTS).
	// +kubebuilder:validation:Optional
	Hsts *bool `json:"hsts,omitempty" tf:"hsts,omitempty"`

	// HSTS includeSubDomains directive.
	// +kubebuilder:validation:Optional
	HstsIncludeSubdomains *bool `json:"hstsIncludeSubdomains,omitempty" tf:"hsts_include_subdomains,omitempty"`

	// HSTS max-age.
	// +kubebuilder:validation:Optional
	HstsMaxAge *int64 `json:"hstsMaxAge,omitempty" tf:"hsts_max_age,omitempty"`

	// HSTS preload directive.
	// +kubebuilder:validation:Optional
	HstsPreload *bool `json:"hstsPreload,omitempty" tf:"hsts_preload,omitempty"`

	// Allow connections via IPv6 to your systems.
	// +kubebuilder:validation:Optional
	IPv6Active *bool `json:"ipv6Active,omitempty" tf:"ipv6_active,omitempty"`

	// If activated, no-cache headers (Cache-Control: [private|no-store|no-cache]) will be ignored.
	// +kubebuilder:validation:Optional
	IgnoreNocache *bool `json:"ignoreNocache,omitempty" tf:"ignore_nocache,omitempty"`

	// Optimization of images.
	// +kubebuilder:validation:Optional
	ImageOptimization *bool `json:"imageOptimization,omitempty" tf:"image_optimization,omitempty"`

	// Not selected HTTP methods will be blocked.
	// +kubebuilder:validation:Optional
	LimitAllowedHTTPMethod []*string `json:"limitAllowedHttpMethod,omitempty" tf:"limit_allowed_http_method,omitempty"`

	// Only selected TLS versions will be used.
	// +kubebuilder:validation:Optional
	LimitTLSVersion []*string `json:"limitTlsVersion,omitempty" tf:"limit_tls_version,omitempty"`

	// Use a different log format.
	// +kubebuilder:validation:Optional
	LogFormat *string `json:"logFormat,omitempty" tf:"log_format,omitempty"`

	// Errors per minute that must occur until a report is sent.
	// +kubebuilder:validation:Optional
	MonitoringAlertThreshold *int64 `json:"monitoringAlertThreshold,omitempty" tf:"monitoring_alert_threshold,omitempty"`

	// Email addresses, to which monitoring emails should be send. Multiple addresses are separated with a space.
	// +kubebuilder:validation:Optional
	MonitoringContactEmail *string `json:"monitoringContactEmail,omitempty" tf:"monitoring_contact_email,omitempty"`

	// Enables / disables the upstream error reporting.
	// +kubebuilder:validation:Optional
	MonitoringSendAlert *bool `json:"monitoringSendAlert,omitempty" tf:"monitoring_send_alert,omitempty"`

	// Activates the X-Myra-SSL Header.
	// +kubebuilder:validation:Optional
	MyraSSLHeader *bool `json:"myraSslHeader,omitempty" tf:"myra_ssl_header,omitempty"`

	// Specifies the error that mark the current upstream as "down".
	// +kubebuilder:validation:Optional
	NextUpstream []*string `json:"nextUpstream,omitempty" tf:"next_upstream,omitempty"`

	// Shall the origin server always be requested via HTTPS?
	// +kubebuilder:validation:Optional
	OnlyHTTPS *bool `json:"onlyHttps,omitempty" tf:"only_https,omitempty"`

	// Connection header.
	// +kubebuilder:validation:Optional
	OriginConnectionHeader *string `json:"originConnectionHeader,omitempty" tf:"origin_connection_header,omitempty"`

	// Name of the cookie which forces Myra to deliver the response not from cache.
	// +kubebuilder:validation:Optional
	ProxyCacheBypass *string `json:"proxyCacheBypass,omitempty" tf:"proxy_cache_bypass,omitempty"`

	// Determines in which cases a stale cached response can be used when an error occurs.
	// +kubebuilder:validation:Optional
	ProxyCacheStale []*string `json:"proxyCacheStale,omitempty" tf:"proxy_cache_stale,omitempty"`

	// Timeout for establishing a connection to the upstream server.
	// +kubebuilder:validation:Optional
	ProxyConnectTimeout *int64 `json:"proxyConnectTimeout,omitempty" tf:"proxy_connect_timeout,omitempty"`

	// Timeout for reading the upstream response.
	// +kubebuilder:validation:Optional
	ProxyReadTimeout *int64 `json:"proxyReadTimeout,omitempty" tf:"proxy_read_timeout,omitempty"`

	// Show CAPTCHA after exceeding the configured request limit.
	// +kubebuilder:validation:Optional
	RequestLimitBlock *string `json:"requestLimitBlock,omitempty" tf:"request_limit_block,omitempty"`

	// Sets how many requests are allowed from an IP per minute.
	// +kubebuilder:validation:Optional
	RequestLimitLevel *int64 `json:"requestLimitLevel,omitempty" tf:"request_limit_level,omitempty"`

	// If activated, an email will be send containing blocked ip addresses that exceeded the configured request limit.
	// +kubebuilder:validation:Optional
	RequestLimitReport *bool `json:"requestLimitReport,omitempty" tf:"request_limit_report,omitempty"`

	// Email addresses, to which request limit emails should be send. Multiple addresses are separated with a space.
	// +kubebuilder:validation:Optional
	RequestLimitReportEmail *string `json:"requestLimitReportEmail,omitempty" tf:"request_limit_report_email,omitempty"`

	// Enable the JavaScript optimization.
	// +kubebuilder:validation:Optional
	Rewrite *bool `json:"rewrite,omitempty" tf:"rewrite,omitempty"`

	// Allows to set a port for communication with origin via SSL.
	// +kubebuilder:validation:Optional
	SSLOriginPort *int64 `json:"sslOriginPort,omitempty" tf:"ssl_origin_port,omitempty"`

	// Protocol to query the origin server.
	// +kubebuilder:validation:Optional
	SourceProtocol *string `json:"sourceProtocol,omitempty" tf:"source_protocol,omitempty"`

	// Activates the SPDY protocol.
	// +kubebuilder:validation:Optional
	Spdy *bool `json:"spdy,omitempty" tf:"spdy,omitempty"`

	// The Subdomain for the Settings.
	// +kubebuilder:validation:Required
	SubdomainName *string `json:"subdomainName" tf:"subdomain_name,omitempty"`

	// Enables / disables the Web Application Firewall.
	// +kubebuilder:validation:Optional
	WafEnable *bool `json:"wafEnable,omitempty" tf:"waf_enable,omitempty"`

	// Level of applied WAF rules.
	// +kubebuilder:validation:Optional
	WafLevelsEnable []*string `json:"wafLevelsEnable,omitempty" tf:"waf_levels_enable,omitempty"`

	// Default policy for the Web Application Firewall in case of rule error.
	// +kubebuilder:validation:Optional
	WafPolicy *string `json:"wafPolicy,omitempty" tf:"waf_policy,omitempty"`
}

// SettingsSpec defines the desired state of Settings
type SettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SettingsParameters `json:"forProvider"`
}

// SettingsStatus defines the observed state of Settings.
type SettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Settings is the Schema for the Settingss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,myrasecjet}
type Settings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SettingsSpec   `json:"spec"`
	Status            SettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SettingsList contains a list of Settingss
type SettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Settings `json:"items"`
}

// Repository type metadata.
var (
	Settings_Kind             = "Settings"
	Settings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Settings_Kind}.String()
	Settings_KindAPIVersion   = Settings_Kind + "." + CRDGroupVersion.String()
	Settings_GroupVersionKind = CRDGroupVersion.WithKind(Settings_Kind)
)

func init() {
	SchemeBuilder.Register(&Settings{}, &SettingsList{})
}
