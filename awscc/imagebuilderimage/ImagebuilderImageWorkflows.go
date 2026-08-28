// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimage


type ImagebuilderImageWorkflows struct {
	// Define execution decision in case of workflow failure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image#on_failure ImagebuilderImage#on_failure}
	OnFailure *string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The parallel group name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image#parallel_group ImagebuilderImage#parallel_group}
	ParallelGroup *string `field:"optional" json:"parallelGroup" yaml:"parallelGroup"`
	// The parameters associated with the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image#parameters ImagebuilderImage#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The Amazon Resource Name (ARN) of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image#workflow_arn ImagebuilderImage#workflow_arn}
	WorkflowArn *string `field:"optional" json:"workflowArn" yaml:"workflowArn"`
}

