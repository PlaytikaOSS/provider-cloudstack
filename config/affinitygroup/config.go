//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package affinitygroup

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_affinity_group", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "AffinityGroup"
	})
}
