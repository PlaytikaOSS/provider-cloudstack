/*
Copyright 2022 Upbound Inc.
*/

package network

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_network", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.References["zone"] = config.Reference{
			TerraformName: "cloudstack_zone",
		}
		r.References["network_offering"] = config.Reference{
			TerraformName: "cloudstack_network_offering",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName: "cloudstack_vpc",
		}
		r.References["acl_id"] = config.Reference{
			TerraformName: "cloudstack_network_acl",
		}
	})
}
