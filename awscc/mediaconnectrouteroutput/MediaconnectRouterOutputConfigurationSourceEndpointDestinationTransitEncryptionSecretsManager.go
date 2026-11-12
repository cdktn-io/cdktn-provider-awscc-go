// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManager struct {
	// The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_output#encryption_secret MediaconnectRouterOutput#encryption_secret}
	EncryptionSecret *MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerEncryptionSecret `field:"optional" json:"encryptionSecret" yaml:"encryptionSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_output#encryption_secret_type MediaconnectRouterOutput#encryption_secret_type}.
	EncryptionSecretType *string `field:"optional" json:"encryptionSecretType" yaml:"encryptionSecretType"`
}

