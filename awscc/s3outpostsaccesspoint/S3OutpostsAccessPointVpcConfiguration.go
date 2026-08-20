// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3outpostsaccesspoint


type S3OutpostsAccessPointVpcConfiguration struct {
	// Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3outposts_access_point#vpc_id S3OutpostsAccessPoint#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

