// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package efsaccesspoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EfsAccessPointConfig struct {
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
	// The ID of the EFS file system that the access point applies to.
	//
	// Accepts only the ID format for input when specifying a file system, for example ``fs-0123456789abcedf2``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/efs_access_point#file_system_id EfsAccessPoint#file_system_id}
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// An array of key-value pairs to apply to this resource.  For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/efs_access_point#access_point_tags EfsAccessPoint#access_point_tags}
	AccessPointTags interface{} `field:"optional" json:"accessPointTags" yaml:"accessPointTags"`
	// The opaque string specified in the request to ensure idempotent creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/efs_access_point#client_token EfsAccessPoint#client_token}
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The full POSIX identity, including the user ID, group ID, and secondary group IDs on the access point that is used for all file operations by NFS clients using the access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/efs_access_point#posix_user EfsAccessPoint#posix_user}
	PosixUser *EfsAccessPointPosixUser `field:"optional" json:"posixUser" yaml:"posixUser"`
	// The directory on the EFS file system that the access point exposes as the root directory to NFS clients using the access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/efs_access_point#root_directory EfsAccessPoint#root_directory}
	RootDirectory *EfsAccessPointRootDirectory `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
}

