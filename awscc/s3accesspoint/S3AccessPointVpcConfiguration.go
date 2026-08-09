// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3accesspoint


type S3AccessPointVpcConfiguration struct {
	// If this field is specified, this access point will only allow connections from the specified VPC ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_access_point#vpc_id S3AccessPoint#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

