/*
Copyright 2022 Upbound Inc.
*/

package volume

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_volume", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "Volume"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudstack_zone",
		}
		r.References["disk_offering_id"] = config.Reference{
			TerraformName: "cloudstack_disk_offering",
		}
	})
}
