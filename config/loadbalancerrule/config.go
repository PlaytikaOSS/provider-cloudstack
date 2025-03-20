//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package loadbalancerrule

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_loadbalancer_rule", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "LoadBalancerRule"
		r.References["ipaddress_id"] = config.Reference{
			TerraformName: "cloudstack_ipaddress",
		}
		r.References["network_id"] = config.Reference{
			TerraformName: "cloudstack_network",
		}
	})
}
