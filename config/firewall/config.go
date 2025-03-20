/*
Copyright 2022 Upbound Inc.
*/

package firewall

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_firewall", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.References["ipaddress_id"] = config.Reference{
			TerraformName: "cloudstack_ipaddress",
		}
	})
}
