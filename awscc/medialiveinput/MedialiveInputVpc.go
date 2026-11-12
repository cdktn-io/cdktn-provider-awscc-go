// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputVpc struct {
	// A list of up to 5 EC2 VPC security group IDs to attach to the Input VPC network interfaces.
	//
	// Requires subnetIds. If none are specified then the VPC default security group will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#security_group_ids MedialiveInput#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of 2 VPC subnet IDs from the same VPC.
	//
	// Subnet IDs must be mapped to two unique availability zones (AZ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#subnet_ids MedialiveInput#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

