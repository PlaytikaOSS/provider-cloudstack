//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package vpncustomergateway

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_vpn_customer_gateway", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "VPNCustomerGateway"
	})
}
