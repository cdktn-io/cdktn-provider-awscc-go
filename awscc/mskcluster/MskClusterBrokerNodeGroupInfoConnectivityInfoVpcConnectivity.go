// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_cluster#client_authentication MskCluster#client_authentication}.
	ClientAuthentication *MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthentication `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
}

