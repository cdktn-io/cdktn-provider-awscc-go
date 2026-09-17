// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowSourceRouterIntegrationTransitDecryption struct {
	// Configuration settings for flow transit encryption keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#encryption_key_configuration MediaconnectFlow#encryption_key_configuration}
	EncryptionKeyConfiguration *MediaconnectFlowSourceRouterIntegrationTransitDecryptionEncryptionKeyConfiguration `field:"optional" json:"encryptionKeyConfiguration" yaml:"encryptionKeyConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#encryption_key_type MediaconnectFlow#encryption_key_type}.
	EncryptionKeyType *string `field:"optional" json:"encryptionKeyType" yaml:"encryptionKeyType"`
}

