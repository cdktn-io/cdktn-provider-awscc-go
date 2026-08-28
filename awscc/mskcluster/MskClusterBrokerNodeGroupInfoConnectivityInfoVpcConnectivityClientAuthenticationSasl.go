// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSasl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_cluster#iam MskCluster#iam}.
	Iam *MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslIam `field:"optional" json:"iam" yaml:"iam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_cluster#scram MskCluster#scram}.
	Scram *MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslScram `field:"optional" json:"scram" yaml:"scram"`
}

