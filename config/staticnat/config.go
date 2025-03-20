//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package staticnat

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_static_nat", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "StaticNAT"
		r.References["ip_address_id"] = config.Reference{
			TerraformName: "cloudstack_ipaddress",
		}
		r.References["virtual_machine_id"] = config.Reference{
			TerraformName: "cloudstack_instance",
		}
	})
}
