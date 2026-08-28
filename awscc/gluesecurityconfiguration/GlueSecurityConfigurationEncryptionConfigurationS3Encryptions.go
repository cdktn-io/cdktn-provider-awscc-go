// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluesecurityconfiguration


type GlueSecurityConfigurationEncryptionConfigurationS3Encryptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_security_configuration#kms_key_arn GlueSecurityConfiguration#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_security_configuration#s3_encryption_mode GlueSecurityConfiguration#s3_encryption_mode}.
	S3EncryptionMode *string `field:"optional" json:"s3EncryptionMode" yaml:"s3EncryptionMode"`
}

