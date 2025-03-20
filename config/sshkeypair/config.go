//nolint:golint
/*
Copyright 2022 Upbound Inc.
*/

package sshkeypair

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_ssh_keypair", func(r *config.Resource) {
		r.ShortGroup = "cloudstack"
		r.Kind = "SSHKeypair"
	})
}
