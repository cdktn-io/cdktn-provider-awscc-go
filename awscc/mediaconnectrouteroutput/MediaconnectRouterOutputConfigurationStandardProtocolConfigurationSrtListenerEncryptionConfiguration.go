// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtListenerEncryptionConfiguration struct {
	// The configuration settings for transit encryption using Secrets Manager, including the secret ARN and role ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_output#encryption_key MediaconnectRouterOutput#encryption_key}
	EncryptionKey *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtListenerEncryptionConfigurationEncryptionKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

