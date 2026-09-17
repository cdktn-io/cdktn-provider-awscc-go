// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationefs

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatasyncLocationEfsConfig struct {
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
	// The subnet and security group that DataSync uses to access target EFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_efs#ec_2_config DatasyncLocationEfs#ec_2_config}
	Ec2Config *DatasyncLocationEfsEc2Config `field:"required" json:"ec2Config" yaml:"ec2Config"`
	// The Amazon Resource Name (ARN) for the Amazon EFS Access point that DataSync uses when accessing the EFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_efs#access_point_arn DatasyncLocationEfs#access_point_arn}
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The Amazon Resource Name (ARN) for the Amazon EFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_efs#efs_filesystem_arn DatasyncLocationEfs#efs_filesystem_arn}
	EfsFilesystemArn *string `field:"optional" json:"efsFilesystemArn" yaml:"efsFilesystemArn"`
	// The Amazon Resource Name (ARN) of the AWS IAM role that the DataSync will assume when mounting the EFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_efs#file_system_access_role_arn DatasyncLocationEfs#file_system_access_role_arn}
	FileSystemAccessRoleArn *string `field:"optional" json:"fileSystemAccessRoleArn" yaml:"fileSystemAccessRoleArn"`
	// Protocol that is used for encrypting the traffic exchanged between the DataSync Agent and the EFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_efs#in_transit_encryption DatasyncLocationEfs#in_transit_encryption}
	InTransitEncryption *string `field:"optional" json:"inTransitEncryption" yaml:"inTransitEncryption"`
	// A subdirectory in the location's path.
	//
	// This subdirectory in the EFS file system is used to read data from the EFS source location or write data to the EFS destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_efs#subdirectory DatasyncLocationEfs#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_efs#tags DatasyncLocationEfs#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

