// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterRemoteNetworkConfig struct {
	// Network configuration of nodes run on-premises with EKS Hybrid Nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#remote_node_networks EksCluster#remote_node_networks}
	RemoteNodeNetworks interface{} `field:"optional" json:"remoteNodeNetworks" yaml:"remoteNodeNetworks"`
	// Network configuration of pods run on-premises with EKS Hybrid Nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#remote_pod_networks EksCluster#remote_pod_networks}
	RemotePodNetworks interface{} `field:"optional" json:"remotePodNetworks" yaml:"remotePodNetworks"`
}

