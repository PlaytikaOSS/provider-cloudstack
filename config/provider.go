/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	    "github.com/playtika/provider-cloudstack/config/account"
        "github.com/playtika/provider-cloudstack/config/affinitygroup"
        "github.com/playtika/provider-cloudstack/config/autoscalevmprofile"
        "github.com/playtika/provider-cloudstack/config/disk"
        "github.com/playtika/provider-cloudstack/config/diskoffering"
        "github.com/playtika/provider-cloudstack/config/domain"
        "github.com/playtika/provider-cloudstack/config/egressfirewall"
        "github.com/playtika/provider-cloudstack/config/firewall"
        "github.com/playtika/provider-cloudstack/config/instance"
        "github.com/playtika/provider-cloudstack/config/ipaddress"
        "github.com/playtika/provider-cloudstack/config/loadbalancerrule"
        "github.com/playtika/provider-cloudstack/config/network"
        "github.com/playtika/provider-cloudstack/config/networkacl"
        "github.com/playtika/provider-cloudstack/config/networkaclrule"
        "github.com/playtika/provider-cloudstack/config/networkoffering"
        "github.com/playtika/provider-cloudstack/config/nic"
        "github.com/playtika/provider-cloudstack/config/portforward"
        "github.com/playtika/provider-cloudstack/config/privategateway"
        "github.com/playtika/provider-cloudstack/config/secondaryipaddress"
        "github.com/playtika/provider-cloudstack/config/securitygroup"
        "github.com/playtika/provider-cloudstack/config/securitygrouprule"
        "github.com/playtika/provider-cloudstack/config/serviceoffering"
        "github.com/playtika/provider-cloudstack/config/sshkeypair"
        "github.com/playtika/provider-cloudstack/config/staticnat"
        "github.com/playtika/provider-cloudstack/config/staticroute"
        "github.com/playtika/provider-cloudstack/config/template"
        "github.com/playtika/provider-cloudstack/config/user"
        "github.com/playtika/provider-cloudstack/config/volume"
        "github.com/playtika/provider-cloudstack/config/vpc"
        "github.com/playtika/provider-cloudstack/config/vpnconnection"
	    "github.com/playtika/provider-cloudstack/config/vpncustomergateway"
	    "github.com/playtika/provider-cloudstack/config/vpngateway"
	    "github.com/playtika/provider-cloudstack/config/zone"

)

const (
	resourcePrefix = "cloudstack"
	modulePath     = "github.com/playtika/provider-cloudstack"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("playtika.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		        account.Configure,
                affinitygroup.Configure,
                autoscalevmprofile.Configure,
                disk.Configure,
                diskoffering.Configure,
                domain.Configure,
                egressfirewall.Configure,
                firewall.Configure,
                instance.Configure,
                ipaddress.Configure,
                loadbalancerrule.Configure,
                network.Configure,
                networkacl.Configure,
                networkaclrule.Configure,
                networkoffering.Configure,
                nic.Configure,
                portforward.Configure,
                privategateway.Configure,
                secondaryipaddress.Configure,
                securitygroup.Configure,
                securitygrouprule.Configure,
                serviceoffering.Configure,
                sshkeypair.Configure,
                staticnat.Configure,
                staticroute.Configure,
                template.Configure,
                user.Configure,
                volume.Configure,
                vpc.Configure,
                vpnconnection.Configure,
                vpncustomergateway.Configure,
                vpngateway.Configure,
                zone.Configure,
		
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
