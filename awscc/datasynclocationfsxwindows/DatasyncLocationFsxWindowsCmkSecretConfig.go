// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationfsxwindows


type DatasyncLocationFsxWindowsCmkSecretConfig struct {
	// Specifies the ARN for the customer-managed AWS KMS key used to encrypt the secret specified for SecretArn.
	//
	// DataSync provides this key to AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datasync_location_fsx_windows#kms_key_arn DatasyncLocationFsxWindows#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

