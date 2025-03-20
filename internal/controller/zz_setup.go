// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/account"
	affinitygroup "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/affinitygroup"
	autoscalevmprofile "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/autoscalevmprofile"
	disk "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/disk"
	diskoffering "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/diskoffering"
	domain "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/domain"
	egressfirewall "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/egressfirewall"
	firewall "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/firewall"
	instance "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/instance"
	ipaddress "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/ipaddress"
	loadbalancerrule "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/loadbalancerrule"
	network "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/network"
	networkacl "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/networkacl"
	networkaclrule "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/networkaclrule"
	networkoffering "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/networkoffering"
	nic "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/nic"
	portforward "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/portforward"
	privategateway "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/privategateway"
	secondaryipaddress "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/secondaryipaddress"
	securitygroup "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/securitygroup"
	securitygrouprule "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/securitygrouprule"
	serviceoffering "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/serviceoffering"
	sshkeypair "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/sshkeypair"
	staticnat "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/staticnat"
	staticroute "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/staticroute"
	template "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/template"
	user "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/user"
	volume "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/volume"
	vpc "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/vpc"
	vpnconnection "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/vpnconnection"
	vpncustomergateway "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/vpncustomergateway"
	vpngateway "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/vpngateway"
	zone "github.com/playtika/provider-cloudstack/internal/controller/cloudstack/zone"
	providerconfig "github.com/playtika/provider-cloudstack/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		affinitygroup.Setup,
		autoscalevmprofile.Setup,
		disk.Setup,
		diskoffering.Setup,
		domain.Setup,
		egressfirewall.Setup,
		firewall.Setup,
		instance.Setup,
		ipaddress.Setup,
		loadbalancerrule.Setup,
		network.Setup,
		networkacl.Setup,
		networkaclrule.Setup,
		networkoffering.Setup,
		nic.Setup,
		portforward.Setup,
		privategateway.Setup,
		secondaryipaddress.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		serviceoffering.Setup,
		sshkeypair.Setup,
		staticnat.Setup,
		staticroute.Setup,
		template.Setup,
		user.Setup,
		volume.Setup,
		vpc.Setup,
		vpnconnection.Setup,
		vpncustomergateway.Setup,
		vpngateway.Setup,
		zone.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
