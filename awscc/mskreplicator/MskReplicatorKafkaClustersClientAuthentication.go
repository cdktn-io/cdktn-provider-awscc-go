// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClustersClientAuthentication struct {
	// Details for mTLS client authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_replicator#mtls MskReplicator#mtls}
	Mtls *MskReplicatorKafkaClustersClientAuthenticationMtls `field:"optional" json:"mtls" yaml:"mtls"`
	// Details for SASL/SCRAM client authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_replicator#sasl_scram MskReplicator#sasl_scram}
	SaslScram *MskReplicatorKafkaClustersClientAuthenticationSaslScram `field:"optional" json:"saslScram" yaml:"saslScram"`
}

