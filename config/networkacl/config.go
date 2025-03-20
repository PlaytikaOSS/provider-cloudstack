//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package networkacl

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_network_acl", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "NetworkACL"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "cloudstack_vpc",
		}
	})
}
