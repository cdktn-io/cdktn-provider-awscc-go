// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineStagesOnFailureConditionsRules struct {
	// The shell commands to run with your compute action in CodePipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codepipeline_pipeline#commands CodepipelinePipeline#commands}
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// The rule's configuration. These are key-value pairs that specify input values for a rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codepipeline_pipeline#configuration CodepipelinePipeline#configuration}
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codepipeline_pipeline#input_artifacts CodepipelinePipeline#input_artifacts}.
	InputArtifacts interface{} `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// The rule declaration's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codepipeline_pipeline#name CodepipelinePipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The rule declaration's AWS Region, such as us-east-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codepipeline_pipeline#region CodepipelinePipeline#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The ARN of the IAM service role that performs the declared rule.
	//
	// This is assumed through the roleArn for the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codepipeline_pipeline#role_arn CodepipelinePipeline#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Represents information about a rule type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codepipeline_pipeline#rule_type_id CodepipelinePipeline#rule_type_id}
	RuleTypeId *CodepipelinePipelineStagesOnFailureConditionsRulesRuleTypeId `field:"optional" json:"ruleTypeId" yaml:"ruleTypeId"`
}

