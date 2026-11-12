// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphrase struct {
	// The ARN of the Secrets Manager secret that holds the key or passphrase.
	//
	// The secret must be encrypted with a customer managed Amazon Web Services KMS key that is enabled and symmetric; it cannot use the Amazon Web Services managed key aws/secretsmanager. The secret must be in the same Amazon Web Services account and Region as the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#secret_arn MediaconnectDestinationEndpoint#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

