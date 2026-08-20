// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationMediaConnectFlowDestinationTransitEncryptionEncryptionKeyConfigurationSecretsManager struct {
	// The ARN of the IAM role assumed by MediaConnect to access the Secrets Manager secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_output#role_arn MediaconnectRouterOutput#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The ARN of the Secrets Manager secret used for transit encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_output#secret_arn MediaconnectRouterOutput#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

