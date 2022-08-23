package vspherevappentity

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vsphere_vapp_entity", func(r *config.Resource) {

        // we need to override the default group that terrajet generated for
        // this resource, which would be "github"
        r.ShortGroup = "vspherevappentity"
    })
}
