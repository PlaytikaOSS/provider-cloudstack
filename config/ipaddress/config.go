/*
Copyright 2022 Upbound Inc.
*/

package ipaddress

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_ipaddress", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "IPAddress"
		r.References["network_id"] = config.Reference{
			TerraformName: "cloudstack_network",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName: "cloudstack_vpc",
		}
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudstack_zone",
		}
	})
}
