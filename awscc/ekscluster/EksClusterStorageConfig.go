// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterStorageConfig struct {
	// Todo: add description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/eks_cluster#block_storage EksCluster#block_storage}
	BlockStorage *EksClusterStorageConfigBlockStorage `field:"optional" json:"blockStorage" yaml:"blockStorage"`
}

