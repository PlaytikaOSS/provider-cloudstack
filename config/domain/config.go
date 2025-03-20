/*
Copyright 2022 Upbound Inc.
*/

package domain

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_domain", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.References["parent_domain_id"] = config.Reference{
			TerraformName: "cloudstack_domain",
		}
	})
}
