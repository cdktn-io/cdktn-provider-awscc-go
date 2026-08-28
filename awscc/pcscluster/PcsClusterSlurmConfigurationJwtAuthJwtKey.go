// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscluster


type PcsClusterSlurmConfigurationJwtAuthJwtKey struct {
	// The Amazon Resource Name (ARN) of the JWT key secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcs_cluster#secret_arn PcsCluster#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The version of the JWT key secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcs_cluster#secret_version PcsCluster#secret_version}
	SecretVersion *string `field:"optional" json:"secretVersion" yaml:"secretVersion"`
}

