// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsSlurmConfig struct {
	// The type of Slurm node for this instance group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#node_type SagemakerCluster#node_type}
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// The Slurm partitions that this instance group belongs to. Maximum of 1 partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#partition_names SagemakerCluster#partition_names}
	PartitionNames *[]*string `field:"optional" json:"partitionNames" yaml:"partitionNames"`
}

