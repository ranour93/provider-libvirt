// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	disk "github.com/nourspeed/provider-libvirt/internal/controller/cloudinit/disk"
	domain "github.com/nourspeed/provider-libvirt/internal/controller/domain/domain"
	network "github.com/nourspeed/provider-libvirt/internal/controller/network/network"
	pool "github.com/nourspeed/provider-libvirt/internal/controller/pool/pool"
	providerconfig "github.com/nourspeed/provider-libvirt/internal/controller/providerconfig"
	volume "github.com/nourspeed/provider-libvirt/internal/controller/volume/volume"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		disk.Setup,
		domain.Setup,
		network.Setup,
		pool.Setup,
		providerconfig.Setup,
		volume.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
