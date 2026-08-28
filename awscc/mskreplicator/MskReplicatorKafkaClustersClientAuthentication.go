// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClustersClientAuthentication struct {
	// Details for mTLS client authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_replicator#mtls MskReplicator#mtls}
	Mtls *MskReplicatorKafkaClustersClientAuthenticationMtls `field:"optional" json:"mtls" yaml:"mtls"`
	// Details for client authentication using SASL/OAUTHBEARER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_replicator#sasl_o_auth_bearer MskReplicator#sasl_o_auth_bearer}
	SaslOAuthBearer *MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearer `field:"optional" json:"saslOAuthBearer" yaml:"saslOAuthBearer"`
	// Details for SASL/SCRAM client authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_replicator#sasl_scram MskReplicator#sasl_scram}
	SaslScram *MskReplicatorKafkaClustersClientAuthenticationSaslScram `field:"optional" json:"saslScram" yaml:"saslScram"`
}

