/*
Copyright 2022 Upbound Inc.
*/

package nic

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_nic", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "NIC"
		r.References["network_id"] = config.Reference{
			TerraformName: "cloudstack_network",
		}
		r.References["virtual_machine_id"] = config.Reference{
			TerraformName: "cloudstack_instance",
		}
	})
}
