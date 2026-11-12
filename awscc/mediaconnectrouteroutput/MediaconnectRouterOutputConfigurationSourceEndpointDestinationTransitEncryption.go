// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryption struct {
	// The configuration for automatic endpoint encryption.
	//
	// In this mode, MediaConnect determines the encryption key when the endpoint connects. This type does not require any additional settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_output#automatic MediaconnectRouterOutput#automatic}
	Automatic *string `field:"optional" json:"automatic" yaml:"automatic"`
	// The endpoint encryption configuration for connections that use a key or passphrase that you store in Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_output#secrets_manager MediaconnectRouterOutput#secrets_manager}
	SecretsManager *MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManager `field:"optional" json:"secretsManager" yaml:"secretsManager"`
}

