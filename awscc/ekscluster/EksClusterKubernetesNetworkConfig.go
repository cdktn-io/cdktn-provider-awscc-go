// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubernetesNetworkConfig struct {
	// Todo: add description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#elastic_load_balancing EksCluster#elastic_load_balancing}
	ElasticLoadBalancing *EksClusterKubernetesNetworkConfigElasticLoadBalancing `field:"optional" json:"elasticLoadBalancing" yaml:"elasticLoadBalancing"`
	// Ipv4 or Ipv6.
	//
	// You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#ip_family EksCluster#ip_family}
	IpFamily *string `field:"optional" json:"ipFamily" yaml:"ipFamily"`
	// The CIDR block to assign Kubernetes service IP addresses from.
	//
	// If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#service_ipv_4_cidr EksCluster#service_ipv_4_cidr}
	ServiceIpv4Cidr *string `field:"optional" json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
}

