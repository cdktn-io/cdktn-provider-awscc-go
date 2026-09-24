// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluemltransform


type GlueMlTransformTransformEncryptionMlUserDataEncryption struct {
	// The ID for the customer-provided KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_ml_transform#kms_key_id GlueMlTransform#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The encryption mode applied to user data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_ml_transform#ml_user_data_encryption_mode GlueMlTransform#ml_user_data_encryption_mode}
	MlUserDataEncryptionMode *string `field:"optional" json:"mlUserDataEncryptionMode" yaml:"mlUserDataEncryptionMode"`
}

