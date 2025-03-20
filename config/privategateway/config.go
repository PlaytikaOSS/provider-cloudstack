//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package privategateway

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_private_gateway", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "PrivateGateway"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "cloudstack_vpc",
		}
		r.References["acl_id"] = config.Reference{
			TerraformName: "cloudstack_network_acl",
		}
		r.References["network_offering"] = config.Reference{
			TerraformName: "cloudstack_network_offering",
		}
		r.References["physical_network_id"] = config.Reference{
			TerraformName: "cloudstack_network",
		}
	})
}
