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

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherecomputecluster"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherecomputeclusterhostgroup"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherecomputeclustervmaffinityrule"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherecomputeclustervmantiaffinityrule"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherecomputeclustervmdependencyrule"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherecomputeclustervmgroup"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherecomputeclustervmhostrule"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherecontentlibrary"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherecontentlibraryitem"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherecustomattribute"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspheredatacenter"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspheredatastorecluster"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspheredatastoreclustervmantiaffinityrule"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspheredistributedportgroup"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspheredistributedvirtualswitch"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspheredrsvmoverride"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspheredpmhostoverride"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherefile"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherefolder"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherehavmoverride"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherehostportgroup"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherehostvirtualswitch"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherelicense"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vsphereresourcepool"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspheretag"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspheretagcategory"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherevirtualdisk"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherevirtualmachine"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherenasdatastore"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherestoragedrsvmoverride"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherevappcontainer"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherevappentity"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherevmfsdatastore"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherevirtualmachinesnapshot"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherehost"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherevnic"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vspherevmstoragepolicy"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vsphererole"
	"github.com/crossplane-contrib/provider-jet-vsphere/config/vsphereentitypermissions"
)

const (
	resourcePrefix = "vsphere"
	modulePath     = "github.com/crossplane-contrib/provider-jet-vsphere"
)

//go:embed schema.json
var providerSchema string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn))//,
		// tjconfig.WithIncludeList([]string{
		// 	"vsphere_compute_cluster$",
		// 	"vsphere_compute_cluster_host_group$",
		// 	"vsphere_compute_cluster_vm_affinity_rule$",
		// 	"vsphere_compute_cluster_vm_anti_affinity_rule$",
		// 	"vsphere_compute_cluster_vm_dependency_rule$",
		// 	"vsphere_compute_cluster_vm_group$",
		// 	"vsphere_compute_cluster_vm_host_rule$",
		// 	"vsphere_content_library$",
		// 	"vsphere_content_library_item$",
		// 	"vsphere_custom_attribute$",
		// 	"vsphere_datacenter$",
		// 	"vsphere_datastore_cluster$",
		// 	"vsphere_datastore_cluster_vm_anti_affinity_rule$",
		// 	"vsphere_distributed_port_group$",
		// 	"vsphere_distributed_virtual_switch$",
		// 	"vsphere_drs_vm_override$",
		// 	"vsphere_dpm_host_override$",
		// 	"vsphere_file$",
		// 	"vsphere_folder$",
		// 	"vsphere_ha_vm_override$",
		// 	"vsphere_host_port_group$",
		// 	"vsphere_host_virtual_switch$",
		// 	"vsphere_license$",
		// 	"vsphere_resource_pool$",
		// 	"vsphere_tag$",
		// 	"vsphere_tag_category$",
		// 	"vsphere_virtual_disk$",
		// 	"vsphere_virtual_machine$",
		// 	"vsphere_nas_datastore$",
		// 	"vsphere_storage_drs_vm_override$",
		// 	"vsphere_vapp_container$",
		// 	"vsphere_vapp_entity$",
		// 	"vsphere_vmfs_datastore$",
		// 	"vsphere_virtual_machine_snapshot$",
		// 	"vsphere_host$",
		// 	"vsphere_vnic$",
		// 	"vsphere_vm_storage_policy$",
		// 	"vsphere_role$",
		// 	"vsphere_entity_permissions$",
		// }))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		vspherecomputecluster.Configure,
		vspherecomputeclusterhostgroup.Configure,
		vspherecomputeclustervmaffinityrule.Configure,
		vspherecomputeclustervmantiaffinityrule.Configure,
		vspherecomputeclustervmdependencyrule.Configure,
		vspherecomputeclustervmgroup.Configure,
		vspherecomputeclustervmhostrule.Configure,
		vspherecontentlibrary.Configure,
		vspherecontentlibraryitem.Configure,
		vspherecustomattribute.Configure,
		vspheredatacenter.Configure,
		vspheredatastorecluster.Configure,
		vspheredatastoreclustervmantiaffinityrule.Configure,
		vspheredistributedportgroup.Configure,
		vspheredistributedvirtualswitch.Configure,
		vspheredrsvmoverride.Configure,
		vspheredpmhostoverride.Configure,
		vspherefile.Configure,
		vspherefolder.Configure,
		vspherehavmoverride.Configure,
		vspherehostportgroup.Configure,
		vspherehostvirtualswitch.Configure,
		vspherelicense.Configure,
		vsphereresourcepool.Configure,
		vspheretag.Configure,
		vspheretagcategory.Configure,
		vspherevirtualdisk.Configure,
		vspherevirtualmachine.Configure,
		vspherenasdatastore.Configure,
		vspherestoragedrsvmoverride.Configure,
		vspherevappcontainer.Configure,
		vspherevappentity.Configure,
		vspherevmfsdatastore.Configure,
		vspherevirtualmachinesnapshot.Configure,
		vspherehost.Configure,
		vspherevnic.Configure,
		vspherevmstoragepolicy.Configure,
		vsphererole.Configure,
		vsphereentitypermissions.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
