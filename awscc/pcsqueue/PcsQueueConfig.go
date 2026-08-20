// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcsqueue

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PcsQueueConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the cluster of the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcs_queue#cluster_id PcsQueue#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The list of compute node group configurations associated with the queue. Queues assign jobs to associated compute node groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcs_queue#compute_node_group_configurations PcsQueue#compute_node_group_configurations}
	ComputeNodeGroupConfigurations interface{} `field:"optional" json:"computeNodeGroupConfigurations" yaml:"computeNodeGroupConfigurations"`
	// The name that identifies the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcs_queue#name PcsQueue#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Slurm configuration for the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcs_queue#slurm_configuration PcsQueue#slurm_configuration}
	SlurmConfiguration *PcsQueueSlurmConfiguration `field:"optional" json:"slurmConfiguration" yaml:"slurmConfiguration"`
	// 1 or more tags added to the resource.
	//
	// Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcs_queue#tags PcsQueue#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

