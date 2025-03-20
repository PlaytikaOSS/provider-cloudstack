//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package secondaryipaddress

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_secondary_ipaddress", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "SecondaryIPAddress"
		r.References["nic_id"] = config.Reference{
			TerraformName: "cloudstack_nic",
		}
		r.References["virtual_machine_id"] = config.Reference{
			TerraformName: "cloudstack_instance",
		}
	})
}
