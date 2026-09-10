// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmassociation


type SsmAssociationOutputLocationS3Location struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_association#output_s3_bucket_name SsmAssociation#output_s3_bucket_name}.
	OutputS3BucketName *string `field:"optional" json:"outputS3BucketName" yaml:"outputS3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_association#output_s3_key_prefix SsmAssociation#output_s3_key_prefix}.
	OutputS3KeyPrefix *string `field:"optional" json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_association#output_s3_region SsmAssociation#output_s3_region}.
	OutputS3Region *string `field:"optional" json:"outputS3Region" yaml:"outputS3Region"`
}

