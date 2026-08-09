// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineStagesActions struct {
	// Represents information about an action type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#action_type_id CodepipelinePipeline#action_type_id}
	ActionTypeId *CodepipelinePipelineStagesActionsActionTypeId `field:"required" json:"actionTypeId" yaml:"actionTypeId"`
	// The action declaration's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#name CodepipelinePipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The shell commands to run with your compute action in CodePipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#commands CodepipelinePipeline#commands}
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// The action's configuration. These are key-value pairs that specify input values for an action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#configuration CodepipelinePipeline#configuration}
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// The list of environment variables that are input to a compute based action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#environment_variables CodepipelinePipeline#environment_variables}
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#input_artifacts CodepipelinePipeline#input_artifacts}.
	InputArtifacts interface{} `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// The variable namespace associated with the action. All variables produced as output by this action fall under this namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#namespace CodepipelinePipeline#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#output_artifacts CodepipelinePipeline#output_artifacts}.
	OutputArtifacts interface{} `field:"optional" json:"outputArtifacts" yaml:"outputArtifacts"`
	// The list of variables that are to be exported from the compute action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#output_variables CodepipelinePipeline#output_variables}
	OutputVariables *[]*string `field:"optional" json:"outputVariables" yaml:"outputVariables"`
	// The action declaration's AWS Region, such as us-east-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#region CodepipelinePipeline#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The ARN of the IAM service role that performs the declared action.
	//
	// This is assumed through the roleArn for the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#role_arn CodepipelinePipeline#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The order in which actions are run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#run_order CodepipelinePipeline#run_order}
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// A timeout duration in minutes that can be applied against the ActionType?s default timeout value specified in Quotas for AWS CodePipeline. This attribute is available only to the manual approval ActionType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_pipeline#timeout_in_minutes CodepipelinePipeline#timeout_in_minutes}
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

