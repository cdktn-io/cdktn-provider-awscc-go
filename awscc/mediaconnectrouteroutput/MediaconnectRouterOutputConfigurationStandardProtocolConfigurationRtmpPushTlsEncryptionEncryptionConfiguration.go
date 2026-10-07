// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushTlsEncryptionEncryptionConfiguration struct {
	// The TLS encryption configuration for destinations that present a certificate from a publicly trusted certificate authority.
	//
	// This type does not require any additional settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_router_output#public MediaconnectRouterOutput#public}
	Public *string `field:"optional" json:"public" yaml:"public"`
}

