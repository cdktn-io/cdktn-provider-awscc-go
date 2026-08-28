// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationFailoverProtocolConfigurationsSrtListenerDecryptionConfiguration struct {
	// The configuration settings for transit encryption using Secrets Manager, including the secret ARN and role ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#encryption_key MediaconnectRouterInput#encryption_key}
	EncryptionKey *MediaconnectRouterInputConfigurationFailoverProtocolConfigurationsSrtListenerDecryptionConfigurationEncryptionKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

