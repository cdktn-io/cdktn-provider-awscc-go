// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputTransitEncryption struct {
	// Defines the configuration settings for transit encryption keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_input#encryption_key_configuration MediaconnectRouterInput#encryption_key_configuration}
	EncryptionKeyConfiguration *MediaconnectRouterInputTransitEncryptionEncryptionKeyConfiguration `field:"optional" json:"encryptionKeyConfiguration" yaml:"encryptionKeyConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_input#encryption_key_type MediaconnectRouterInput#encryption_key_type}.
	EncryptionKeyType *string `field:"optional" json:"encryptionKeyType" yaml:"encryptionKeyType"`
}

