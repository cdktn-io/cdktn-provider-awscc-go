// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationlibrary


type BedrockDataAutomationLibraryEncryptionConfiguration struct {
	// KMS Encryption Context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_automation_library#kms_encryption_context BedrockDataAutomationLibrary#kms_encryption_context}
	KmsEncryptionContext *map[string]*string `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	// KMS Key Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_automation_library#kms_key_id BedrockDataAutomationLibrary#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

