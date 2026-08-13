// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationMediaLiveChannelSourceTransitDecryption struct {
	// Configuration settings for the MediaLive transit encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_input#encryption_key_configuration MediaconnectRouterInput#encryption_key_configuration}
	EncryptionKeyConfiguration *MediaconnectRouterInputConfigurationMediaLiveChannelSourceTransitDecryptionEncryptionKeyConfiguration `field:"optional" json:"encryptionKeyConfiguration" yaml:"encryptionKeyConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_input#encryption_key_type MediaconnectRouterInput#encryption_key_type}.
	EncryptionKeyType *string `field:"optional" json:"encryptionKeyType" yaml:"encryptionKeyType"`
}

