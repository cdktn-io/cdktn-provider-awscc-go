// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3accessgrant


type S3AccessGrantAccessGrantsLocationConfiguration struct {
	// The S3 sub prefix of a registered location in your S3 Access Grants instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_access_grant#s3_sub_prefix S3AccessGrant#s3_sub_prefix}
	S3SubPrefix *string `field:"optional" json:"s3SubPrefix" yaml:"s3SubPrefix"`
}

