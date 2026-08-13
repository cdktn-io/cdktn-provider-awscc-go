// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentS3AccessPoint struct {
	// The S3 access point's policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/fsx_s3_access_point_attachment#policy FsxS3AccessPointAttachment#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// The S3 access point's virtual private cloud (VPC) configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/fsx_s3_access_point_attachment#vpc_configuration FsxS3AccessPointAttachment#vpc_configuration}
	VpcConfiguration *FsxS3AccessPointAttachmentS3AccessPointVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

