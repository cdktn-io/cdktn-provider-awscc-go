// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentS3AccessPointVpcConfiguration struct {
	// Specifies the virtual private cloud (VPC) for the S3 access point VPC configuration, if one exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/fsx_s3_access_point_attachment#vpc_id FsxS3AccessPointAttachment#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

