package account

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_account", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
	})
}
