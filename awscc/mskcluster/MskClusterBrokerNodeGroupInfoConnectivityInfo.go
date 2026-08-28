// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterBrokerNodeGroupInfoConnectivityInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_cluster#network_type MskCluster#network_type}.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_cluster#public_access MskCluster#public_access}.
	PublicAccess *MskClusterBrokerNodeGroupInfoConnectivityInfoPublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_cluster#vpc_connectivity MskCluster#vpc_connectivity}.
	VpcConnectivity *MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivity `field:"optional" json:"vpcConnectivity" yaml:"vpcConnectivity"`
}

