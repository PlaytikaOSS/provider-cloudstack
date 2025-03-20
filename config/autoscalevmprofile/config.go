//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package autoscalevmprofile

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_autoscale_vm_profile", func(r *config.Resource) {
		r.Kind = "AutoscaleVmProfile"
		r.ShortGroup = "cloudstack"
		r.References["template"] = config.Reference{
			TerraformName: "cloudstack_template",
		}
		r.References["service_offering"] = config.Reference{
			TerraformName: "cloudstack_service_offering",
		}
		r.References["zone"] = config.Reference{
			TerraformName: "cloudstack_zone",
		}
	})
}
