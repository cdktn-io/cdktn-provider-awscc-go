// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisepipeline


type IotsitewisePipelineComputations struct {
	// The unique name for this compute node within the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_pipeline#compute_node_name IotsitewisePipeline#compute_node_name}
	ComputeNodeName *string `field:"required" json:"computeNodeName" yaml:"computeNodeName"`
	// The name of the task to execute for this compute node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_pipeline#task_name IotsitewisePipeline#task_name}
	TaskName *string `field:"required" json:"taskName" yaml:"taskName"`
	// A list of compute node names that must complete successfully before this node can start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_pipeline#depends_on IotsitewisePipeline#depends_on}
	DependsOn *[]*string `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// A map of environment variable key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_pipeline#environment_variables IotsitewisePipeline#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
}

