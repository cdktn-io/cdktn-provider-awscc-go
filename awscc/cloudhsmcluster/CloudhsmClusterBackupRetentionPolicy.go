// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudhsmcluster


type CloudhsmClusterBackupRetentionPolicy struct {
	// The type of backup retention policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudhsm_cluster#type CloudhsmCluster#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Use a value between 7 - 379.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudhsm_cluster#value CloudhsmCluster#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

