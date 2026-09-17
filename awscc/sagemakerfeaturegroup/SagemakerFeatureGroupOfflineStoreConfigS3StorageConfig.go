// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup


type SagemakerFeatureGroupOfflineStoreConfigS3StorageConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_feature_group#kms_key_id SagemakerFeatureGroup#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_feature_group#s3_uri SagemakerFeatureGroup#s3_uri}.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

