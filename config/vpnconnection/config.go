//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package vpnconnection

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_vpn_connection", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "VPNConnection"
		r.References["customer_gateway_id"] = config.Reference{
			TerraformName: "cloudstack_vpn_customer_gateway",
		}
		r.References["vpn_gateway_id"] = config.Reference{
			TerraformName: "cloudstack_vpn_gateway",
		}
	})
}
