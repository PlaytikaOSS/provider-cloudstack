/*
Copyright 2022 Upbound Inc.
*/

package vpc

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_vpc", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "VPC"
		r.References["zone"] = config.Reference{
			TerraformName: "cloudstack_zone",
		}
	})
}
