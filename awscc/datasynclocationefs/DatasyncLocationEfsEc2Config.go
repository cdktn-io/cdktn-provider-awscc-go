// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationefs


type DatasyncLocationEfsEc2Config struct {
	// The Amazon Resource Names (ARNs) of the security groups that are configured for the Amazon EC2 resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_efs#security_group_arns DatasyncLocationEfs#security_group_arns}
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// The ARN of the subnet that DataSync uses to access the target EFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_efs#subnet_arn DatasyncLocationEfs#subnet_arn}
	SubnetArn *string `field:"required" json:"subnetArn" yaml:"subnetArn"`
}

