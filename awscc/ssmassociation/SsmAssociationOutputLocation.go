// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmassociation


type SsmAssociationOutputLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_association#s3_location SsmAssociation#s3_location}.
	S3Location *SsmAssociationOutputLocationS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

