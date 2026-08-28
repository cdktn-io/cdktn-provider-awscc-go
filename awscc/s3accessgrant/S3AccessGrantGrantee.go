// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3accessgrant


type S3AccessGrantGrantee struct {
	// The unique identifier of the Grantee.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_access_grant#grantee_identifier S3AccessGrant#grantee_identifier}
	GranteeIdentifier *string `field:"required" json:"granteeIdentifier" yaml:"granteeIdentifier"`
	// Configures the transfer acceleration state for an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_access_grant#grantee_type S3AccessGrant#grantee_type}
	GranteeType *string `field:"required" json:"granteeType" yaml:"granteeType"`
}

