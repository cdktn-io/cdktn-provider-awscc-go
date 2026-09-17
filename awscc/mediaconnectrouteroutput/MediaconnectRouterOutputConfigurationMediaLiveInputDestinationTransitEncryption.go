// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationMediaLiveInputDestinationTransitEncryption struct {
	// Configuration settings for the MediaLive transit encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_router_output#encryption_key_configuration MediaconnectRouterOutput#encryption_key_configuration}
	EncryptionKeyConfiguration *MediaconnectRouterOutputConfigurationMediaLiveInputDestinationTransitEncryptionEncryptionKeyConfiguration `field:"optional" json:"encryptionKeyConfiguration" yaml:"encryptionKeyConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_router_output#encryption_key_type MediaconnectRouterOutput#encryption_key_type}.
	EncryptionKeyType *string `field:"optional" json:"encryptionKeyType" yaml:"encryptionKeyType"`
}

