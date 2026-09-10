// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentials struct {
	// Secrets Manager ARN of the secret containing the client_id and client_secret used to obtain an OAuth Bearer token via the client_credentials grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/msk_replicator#token_request_secret_arn MskReplicator#token_request_secret_arn}
	TokenRequestSecretArn *string `field:"optional" json:"tokenRequestSecretArn" yaml:"tokenRequestSecretArn"`
}

