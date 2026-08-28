// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeApiServerConfig struct {
	// The duration that Kubernetes events are retained (e.g., 30m, 1h).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_cluster#event_ttl EksCluster#event_ttl}
	EventTtl *string `field:"optional" json:"eventTtl" yaml:"eventTtl"`
	// The port range for Kubernetes NodePort services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_cluster#service_node_port_range EksCluster#service_node_port_range}
	ServiceNodePortRange *EksClusterKubeApiServerConfigServiceNodePortRange `field:"optional" json:"serviceNodePortRange" yaml:"serviceNodePortRange"`
}

