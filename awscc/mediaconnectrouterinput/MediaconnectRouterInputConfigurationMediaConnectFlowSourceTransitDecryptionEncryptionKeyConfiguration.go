// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationMediaConnectFlowSourceTransitDecryptionEncryptionKeyConfiguration struct {
	// Configuration settings for automatic encryption key management, where MediaConnect handles key creation and rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_input#automatic MediaconnectRouterInput#automatic}
	Automatic *string `field:"optional" json:"automatic" yaml:"automatic"`
	// The configuration settings for transit encryption using Secrets Manager, including the secret ARN and role ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_input#secrets_manager MediaconnectRouterInput#secrets_manager}
	SecretsManager *MediaconnectRouterInputConfigurationMediaConnectFlowSourceTransitDecryptionEncryptionKeyConfigurationSecretsManager `field:"optional" json:"secretsManager" yaml:"secretsManager"`
}

