// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionmatchingworkflow

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EntityresolutionMatchingWorkflowConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#input_source_config EntityresolutionMatchingWorkflow#input_source_config}.
	InputSourceConfig interface{} `field:"required" json:"inputSourceConfig" yaml:"inputSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#output_source_config EntityresolutionMatchingWorkflow#output_source_config}.
	OutputSourceConfig interface{} `field:"required" json:"outputSourceConfig" yaml:"outputSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#resolution_techniques EntityresolutionMatchingWorkflow#resolution_techniques}.
	ResolutionTechniques *EntityresolutionMatchingWorkflowResolutionTechniques `field:"required" json:"resolutionTechniques" yaml:"resolutionTechniques"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#role_arn EntityresolutionMatchingWorkflow#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the MatchingWorkflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#workflow_name EntityresolutionMatchingWorkflow#workflow_name}
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
	// The description of the MatchingWorkflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#description EntityresolutionMatchingWorkflow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#incremental_run_config EntityresolutionMatchingWorkflow#incremental_run_config}.
	IncrementalRunConfig *EntityresolutionMatchingWorkflowIncrementalRunConfig `field:"optional" json:"incrementalRunConfig" yaml:"incrementalRunConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/entityresolution_matching_workflow#tags EntityresolutionMatchingWorkflow#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

