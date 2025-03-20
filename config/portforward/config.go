//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package portforward

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_port_forward", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "PortForward"
		r.References["ip_address_id"] = config.Reference{
			TerraformName: "cloudstack_ipaddress",
		}
	})
}
