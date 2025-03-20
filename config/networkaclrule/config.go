//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package networkaclrule

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_network_acl_rule", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "NetworkACLRule"
		r.References["acl_id"] = config.Reference{
			TerraformName: "cloudstack_network_acl",
		}
	})
}
