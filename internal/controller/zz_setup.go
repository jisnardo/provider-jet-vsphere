/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	providerconfig "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/providerconfig"
	cluster "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherecomputecluster/cluster"
	clusterhostgroup "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherecomputeclusterhostgroup/clusterhostgroup"
	clustervmaffinityrule "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherecomputeclustervmaffinityrule/clustervmaffinityrule"
	clustervmantiaffinityrule "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherecomputeclustervmantiaffinityrule/clustervmantiaffinityrule"
	clustervmdependencyrule "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherecomputeclustervmdependencyrule/clustervmdependencyrule"
	clustervmgroup "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherecomputeclustervmgroup/clustervmgroup"
	clustervmhostrule "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherecomputeclustervmhostrule/clustervmhostrule"
	library "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherecontentlibrary/library"
	libraryitem "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherecontentlibraryitem/libraryitem"
	attribute "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherecustomattribute/attribute"
	datacenter "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspheredatacenter/datacenter"
	clustervspheredatastorecluster "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspheredatastorecluster/cluster"
	clustervmantiaffinityrulevspheredatastoreclustervmantiaffinityrule "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspheredatastoreclustervmantiaffinityrule/clustervmantiaffinityrule"
	portgroup "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspheredistributedportgroup/portgroup"
	virtualswitch "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspheredistributedvirtualswitch/virtualswitch"
	hostoverride "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspheredpmhostoverride/hostoverride"
	vmoverride "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspheredrsvmoverride/vmoverride"
	permissions "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vsphereentitypermissions/permissions"
	file "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherefile/file"
	folder "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherefolder/folder"
	vmoverridevspherehavmoverride "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherehavmoverride/vmoverride"
	host "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherehost/host"
	portgroupvspherehostportgroup "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherehostportgroup/portgroup"
	virtualswitchvspherehostvirtualswitch "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherehostvirtualswitch/virtualswitch"
	license "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherelicense/license"
	datastore "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherenasdatastore/datastore"
	pool "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vsphereresourcepool/pool"
	role "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vsphererole/role"
	drsvmoverride "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherestoragedrsvmoverride/drsvmoverride"
	tag "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspheretag/tag"
	category "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspheretagcategory/category"
	container "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherevappcontainer/container"
	entity "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherevappentity/entity"
	disk "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherevirtualdisk/disk"
	machine "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherevirtualmachine/machine"
	machinesnapshot "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherevirtualmachinesnapshot/machinesnapshot"
	datastorevspherevmfsdatastore "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherevmfsdatastore/datastore"
	storagepolicy "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherevmstoragepolicy/storagepolicy"
	vnic "github.com/crossplane-contrib/provider-jet-vsphere/internal/controller/vspherevnic/vnic"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		cluster.Setup,
		clusterhostgroup.Setup,
		clustervmaffinityrule.Setup,
		clustervmantiaffinityrule.Setup,
		clustervmdependencyrule.Setup,
		clustervmgroup.Setup,
		clustervmhostrule.Setup,
		library.Setup,
		libraryitem.Setup,
		attribute.Setup,
		datacenter.Setup,
		clustervspheredatastorecluster.Setup,
		clustervmantiaffinityrulevspheredatastoreclustervmantiaffinityrule.Setup,
		portgroup.Setup,
		virtualswitch.Setup,
		hostoverride.Setup,
		vmoverride.Setup,
		permissions.Setup,
		file.Setup,
		folder.Setup,
		vmoverridevspherehavmoverride.Setup,
		host.Setup,
		portgroupvspherehostportgroup.Setup,
		virtualswitchvspherehostvirtualswitch.Setup,
		license.Setup,
		datastore.Setup,
		pool.Setup,
		role.Setup,
		drsvmoverride.Setup,
		tag.Setup,
		category.Setup,
		container.Setup,
		entity.Setup,
		disk.Setup,
		machine.Setup,
		machinesnapshot.Setup,
		datastorevspherevmfsdatastore.Setup,
		storagepolicy.Setup,
		vnic.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
