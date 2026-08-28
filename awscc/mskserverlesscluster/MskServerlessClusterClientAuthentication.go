// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskserverlesscluster


type MskServerlessClusterClientAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_serverless_cluster#sasl MskServerlessCluster#sasl}.
	Sasl *MskServerlessClusterClientAuthenticationSasl `field:"required" json:"sasl" yaml:"sasl"`
}

