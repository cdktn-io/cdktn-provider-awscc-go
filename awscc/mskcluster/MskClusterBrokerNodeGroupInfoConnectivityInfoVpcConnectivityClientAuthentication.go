// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_cluster#sasl MskCluster#sasl}.
	Sasl *MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSasl `field:"optional" json:"sasl" yaml:"sasl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_cluster#tls MskCluster#tls}.
	Tls *MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationTls `field:"optional" json:"tls" yaml:"tls"`
}

