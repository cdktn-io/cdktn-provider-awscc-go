// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterRollbackConfig struct {
	// The timeout in minutes for the version rollback operation. If not specified, defaults to 720 minutes (12 hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/eks_cluster#timeout_minutes EksCluster#timeout_minutes}
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

