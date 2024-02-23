/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/nourspeed/provider-libvirt/config/pool"
	"github.com/nourspeed/provider-libvirt/config/volume"
	"github.com/nourspeed/provider-libvirt/config/cloudinit"
	"github.com/nourspeed/provider-libvirt/config/domain"
	"github.com/nourspeed/provider-libvirt/config/network"
)

const (
	resourcePrefix = "libvirt"
	modulePath     = "github.com/nourspeed/provider-libvirt"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("nourspeed.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		pool.Configure,
		volume.Configure,
		cloudinit.Configure,
		domain.Configure,
		network.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
