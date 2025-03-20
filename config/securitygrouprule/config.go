//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package securitygrouprule

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_security_group_rule", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "SecurityGroupRule"
		r.References["security_group_id"] = config.Reference{
			TerraformName: "cloudstack_security_group",
		}
	})
}
