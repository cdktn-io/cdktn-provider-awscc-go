// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterLoggingClusterLogging struct {
	// Enable control plane logs for your cluster, all log types will be disabled if the array is empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_cluster#enabled_types EksCluster#enabled_types}
	EnabledTypes interface{} `field:"optional" json:"enabledTypes" yaml:"enabledTypes"`
}

