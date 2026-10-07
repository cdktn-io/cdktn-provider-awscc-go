// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistry


type AgentregistryRegistryEncryptionConfiguration struct {
	// The Amazon Resource Name (ARN) of the customer-managed AWS KMS key used to encrypt the registry's content.
	//
	// The key must be a symmetric encryption key in the same AWS account and Region as the registry. Multi-Region keys are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/agentregistry_registry#kms_key_arn AgentregistryRegistry#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

