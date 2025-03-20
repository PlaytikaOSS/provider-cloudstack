//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package staticroute

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_static_route", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "StaticRoute"
		r.References["gateway_id"] = config.Reference{
			TerraformName: "cloudstack_private_gateway",
		}
	})
}
