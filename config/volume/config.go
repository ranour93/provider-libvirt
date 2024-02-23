package volume

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("libvirt_volume", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "libvirt"
        r.ShortGroup = "volume"

        r.References["pool"] = config.Reference{
            Type: "github.com/nourspeed/provider-libvirt/apis/pool/v1alpha1.Pool",
        }
    })
}