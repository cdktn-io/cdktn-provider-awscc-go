// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterRemoteNetworkConfigRemoteNodeNetworks struct {
	// Specifies the list of remote node CIDRs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_cluster#cidrs EksCluster#cidrs}
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
}

