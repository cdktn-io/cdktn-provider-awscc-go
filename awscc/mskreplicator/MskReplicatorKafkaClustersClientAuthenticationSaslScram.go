// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClustersClientAuthenticationSaslScram struct {
	// The SASL/SCRAM authentication mechanism.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_replicator#mechanism MskReplicator#mechanism}
	Mechanism *string `field:"optional" json:"mechanism" yaml:"mechanism"`
	// The Amazon Resource Name (ARN) of the Secrets Manager secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_replicator#secret_arn MskReplicator#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

