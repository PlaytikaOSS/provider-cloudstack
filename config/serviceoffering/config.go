//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package serviceoffering

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_service_offering", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "ServiceOffering"
	})
}
