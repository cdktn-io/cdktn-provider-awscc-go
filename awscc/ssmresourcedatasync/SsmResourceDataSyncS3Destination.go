// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmresourcedatasync


type SsmResourceDataSyncS3Destination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_resource_data_sync#bucket_name SsmResourceDataSync#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_resource_data_sync#bucket_prefix SsmResourceDataSync#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_resource_data_sync#bucket_region SsmResourceDataSync#bucket_region}.
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_resource_data_sync#kms_key_arn SsmResourceDataSync#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_resource_data_sync#sync_format SsmResourceDataSync#sync_format}.
	SyncFormat *string `field:"optional" json:"syncFormat" yaml:"syncFormat"`
}

