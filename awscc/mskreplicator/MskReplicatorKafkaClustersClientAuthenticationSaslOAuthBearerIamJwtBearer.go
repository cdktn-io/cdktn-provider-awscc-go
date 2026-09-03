// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearer struct {
	// The audience (aud claim) set in the STS JWT assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#audience MskReplicator#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// The algorithm used to sign the JWT assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#signing_algorithm MskReplicator#signing_algorithm}
	SigningAlgorithm *string `field:"optional" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// Optional Secrets Manager ARN for identity providers that require client authentication alongside the JWT Bearer assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#token_request_secret_arn MskReplicator#token_request_secret_arn}
	TokenRequestSecretArn *string `field:"optional" json:"tokenRequestSecretArn" yaml:"tokenRequestSecretArn"`
}

