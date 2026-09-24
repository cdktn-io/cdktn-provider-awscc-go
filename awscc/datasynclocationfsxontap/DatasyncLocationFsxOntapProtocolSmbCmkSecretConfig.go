// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationfsxontap


type DatasyncLocationFsxOntapProtocolSmbCmkSecretConfig struct {
	// Specifies the ARN for the customer-managed AWS KMS key used to encrypt the secret specified for SecretArn.
	//
	// DataSync provides this key to AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_fsx_ontap#kms_key_arn DatasyncLocationFsxOntap#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

