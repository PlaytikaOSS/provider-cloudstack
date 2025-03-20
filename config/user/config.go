/*
Copyright 2022 Upbound Inc.
*/

package user

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_user", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "User"
		r.References["account"] = config.Reference{
			TerraformName: "cloudstack_account",
		}
	})
}
