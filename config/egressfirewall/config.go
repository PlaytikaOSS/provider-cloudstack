//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package egressfirewall

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_egress_firewall", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "EgressFirewall"
		r.References["network_id"] = config.Reference{
			TerraformName: "cloudstack_network",
		}
	})
}
