/*
Copyright 2022 Upbound Inc.
*/

package template

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_template", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "Template"
		r.References["zone"] = config.Reference{
			TerraformName: "cloudstack_zone",
		}
	})
}
