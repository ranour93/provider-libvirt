package domain

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("libvirt_domain", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "libvirt"
		r.ShortGroup = "domain"

		r.References["cloudinit"] = config.Reference{
			Type: "github.com/nourspeed/provider-libvirt/apis/cloudinit/v1alpha1.Disk",
		}
	})
}
