// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationnfs

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatasyncLocationNfsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Contains a list of Amazon Resource Names (ARNs) of agents that are used to connect an NFS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_nfs#on_prem_config DatasyncLocationNfs#on_prem_config}
	OnPremConfig *DatasyncLocationNfsOnPremConfig `field:"required" json:"onPremConfig" yaml:"onPremConfig"`
	// The NFS mount options that DataSync can use to mount your NFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_nfs#mount_options DatasyncLocationNfs#mount_options}
	MountOptions *DatasyncLocationNfsMountOptions `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The name of the NFS server. This value is the IP address or DNS name of the NFS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_nfs#server_hostname DatasyncLocationNfs#server_hostname}
	ServerHostname *string `field:"optional" json:"serverHostname" yaml:"serverHostname"`
	// The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_nfs#subdirectory DatasyncLocationNfs#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_nfs#tags DatasyncLocationNfs#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

