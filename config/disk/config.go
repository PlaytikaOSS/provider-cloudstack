/*
Copyright 2022 Upbound Inc.
*/

package disk

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_disk", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.References["disk_offering"] = config.Reference{
			TerraformName: "cloudstack_disk_offering",
		}
		r.References["virtual_machine"] = config.Reference{
			TerraformName: "cloudstack_instance",
		}
		r.References["zone"] = config.Reference{
			TerraformName: "cloudstack_zone",
		}
	})
}
