package instance

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_instance", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
	})
}
