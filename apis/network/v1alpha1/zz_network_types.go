// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DHCPInitParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DHCPObservation struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DHCPParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DNSInitParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	Forwarders []ForwardersInitParameters `json:"forwarders,omitempty" tf:"forwarders,omitempty"`

	Hosts []HostsInitParameters `json:"hosts,omitempty" tf:"hosts,omitempty"`

	LocalOnly *bool `json:"localOnly,omitempty" tf:"local_only,omitempty"`

	Srvs []SrvsInitParameters `json:"srvs,omitempty" tf:"srvs,omitempty"`
}

type DNSObservation struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	Forwarders []ForwardersObservation `json:"forwarders,omitempty" tf:"forwarders,omitempty"`

	Hosts []HostsObservation `json:"hosts,omitempty" tf:"hosts,omitempty"`

	LocalOnly *bool `json:"localOnly,omitempty" tf:"local_only,omitempty"`

	Srvs []SrvsObservation `json:"srvs,omitempty" tf:"srvs,omitempty"`
}

type DNSParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Forwarders []ForwardersParameters `json:"forwarders,omitempty" tf:"forwarders,omitempty"`

	// +kubebuilder:validation:Optional
	Hosts []HostsParameters `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// +kubebuilder:validation:Optional
	LocalOnly *bool `json:"localOnly,omitempty" tf:"local_only,omitempty"`

	// +kubebuilder:validation:Optional
	Srvs []SrvsParameters `json:"srvs,omitempty" tf:"srvs,omitempty"`
}

type DnsmasqOptionsInitParameters struct {
	Options []OptionsInitParameters `json:"options,omitempty" tf:"options,omitempty"`
}

type DnsmasqOptionsObservation struct {
	Options []OptionsObservation `json:"options,omitempty" tf:"options,omitempty"`
}

type DnsmasqOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Options []OptionsParameters `json:"options,omitempty" tf:"options,omitempty"`
}

type ForwardersInitParameters struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`
}

type ForwardersObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`
}

type ForwardersParameters struct {

	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`
}

type HostsInitParameters struct {
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`
}

type HostsObservation struct {
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`
}

type HostsParameters struct {

	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`
}

type NetworkInitParameters struct {
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	Autostart *bool `json:"autostart,omitempty" tf:"autostart,omitempty"`

	Bridge *string `json:"bridge,omitempty" tf:"bridge,omitempty"`

	DHCP []DHCPInitParameters `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	DNS []DNSInitParameters `json:"dns,omitempty" tf:"dns,omitempty"`

	DnsmasqOptions []DnsmasqOptionsInitParameters `json:"dnsmasqOptions,omitempty" tf:"dnsmasq_options,omitempty"`

	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Routes []RoutesInitParameters `json:"routes,omitempty" tf:"routes,omitempty"`

	XML []XMLInitParameters `json:"xml,omitempty" tf:"xml,omitempty"`
}

type NetworkObservation struct {
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	Autostart *bool `json:"autostart,omitempty" tf:"autostart,omitempty"`

	Bridge *string `json:"bridge,omitempty" tf:"bridge,omitempty"`

	DHCP []DHCPObservation `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	DNS []DNSObservation `json:"dns,omitempty" tf:"dns,omitempty"`

	DnsmasqOptions []DnsmasqOptionsObservation `json:"dnsmasqOptions,omitempty" tf:"dnsmasq_options,omitempty"`

	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Routes []RoutesObservation `json:"routes,omitempty" tf:"routes,omitempty"`

	XML []XMLObservation `json:"xml,omitempty" tf:"xml,omitempty"`
}

type NetworkParameters struct {

	// +kubebuilder:validation:Optional
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// +kubebuilder:validation:Optional
	Autostart *bool `json:"autostart,omitempty" tf:"autostart,omitempty"`

	// +kubebuilder:validation:Optional
	Bridge *string `json:"bridge,omitempty" tf:"bridge,omitempty"`

	// +kubebuilder:validation:Optional
	DHCP []DHCPParameters `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// +kubebuilder:validation:Optional
	DNS []DNSParameters `json:"dns,omitempty" tf:"dns,omitempty"`

	// +kubebuilder:validation:Optional
	DnsmasqOptions []DnsmasqOptionsParameters `json:"dnsmasqOptions,omitempty" tf:"dnsmasq_options,omitempty"`

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Routes []RoutesParameters `json:"routes,omitempty" tf:"routes,omitempty"`

	// +kubebuilder:validation:Optional
	XML []XMLParameters `json:"xml,omitempty" tf:"xml,omitempty"`
}

type OptionsInitParameters struct {
	OptionName *string `json:"optionName,omitempty" tf:"option_name,omitempty"`

	OptionValue *string `json:"optionValue,omitempty" tf:"option_value,omitempty"`
}

type OptionsObservation struct {
	OptionName *string `json:"optionName,omitempty" tf:"option_name,omitempty"`

	OptionValue *string `json:"optionValue,omitempty" tf:"option_value,omitempty"`
}

type OptionsParameters struct {

	// +kubebuilder:validation:Optional
	OptionName *string `json:"optionName,omitempty" tf:"option_name,omitempty"`

	// +kubebuilder:validation:Optional
	OptionValue *string `json:"optionValue,omitempty" tf:"option_value,omitempty"`
}

type RoutesInitParameters struct {
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`
}

type RoutesObservation struct {
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`
}

type RoutesParameters struct {

	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr" tf:"cidr,omitempty"`

	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway" tf:"gateway,omitempty"`
}

type SrvsInitParameters struct {
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	Weight *string `json:"weight,omitempty" tf:"weight,omitempty"`
}

type SrvsObservation struct {
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	Weight *string `json:"weight,omitempty" tf:"weight,omitempty"`
}

type SrvsParameters struct {

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// +kubebuilder:validation:Optional
	Weight *string `json:"weight,omitempty" tf:"weight,omitempty"`
}

type XMLInitParameters struct {
	Xslt *string `json:"xslt,omitempty" tf:"xslt,omitempty"`
}

type XMLObservation struct {
	Xslt *string `json:"xslt,omitempty" tf:"xslt,omitempty"`
}

type XMLParameters struct {

	// +kubebuilder:validation:Optional
	Xslt *string `json:"xslt,omitempty" tf:"xslt,omitempty"`
}

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider NetworkInitParameters `json:"initProvider,omitempty"`
}

// NetworkStatus defines the observed state of Network.
type NetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Network is the Schema for the Networks API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,libvirt}
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   NetworkSpec   `json:"spec"`
	Status NetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkList contains a list of Networks
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

// Repository type metadata.
var (
	Network_Kind             = "Network"
	Network_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Network_Kind}.String()
	Network_KindAPIVersion   = Network_Kind + "." + CRDGroupVersion.String()
	Network_GroupVersionKind = CRDGroupVersion.WithKind(Network_Kind)
)

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}
