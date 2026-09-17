// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowoutput


type MediaconnectFlowOutputRouterIntegrationTransitEncryptionEncryptionKeyConfiguration struct {
	// Configuration settings for automatic encryption key management, where MediaConnect handles key creation and rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow_output#automatic MediaconnectFlowOutput#automatic}
	Automatic *string `field:"optional" json:"automatic" yaml:"automatic"`
	// The configuration settings for transit encryption of a flow output using AWS Secrets Manager, including the secret ARN and role ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow_output#secrets_manager MediaconnectFlowOutput#secrets_manager}
	SecretsManager *MediaconnectFlowOutputRouterIntegrationTransitEncryptionEncryptionKeyConfigurationSecretsManager `field:"optional" json:"secretsManager" yaml:"secretsManager"`
}

