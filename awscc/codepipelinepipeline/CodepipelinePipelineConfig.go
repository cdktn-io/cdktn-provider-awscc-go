// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CodepipelinePipelineConfig struct {
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
	// The Amazon Resource Name (ARN) for CodePipeline to use to either perform actions with no actionRoleArn, or to use to assume roles for actions with an actionRoleArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#role_arn CodepipelinePipeline#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Represents information about a stage and its definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#stages CodepipelinePipeline#stages}
	Stages interface{} `field:"required" json:"stages" yaml:"stages"`
	// The S3 bucket where artifacts for the pipeline are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#artifact_store CodepipelinePipeline#artifact_store}
	ArtifactStore *CodepipelinePipelineArtifactStore `field:"optional" json:"artifactStore" yaml:"artifactStore"`
	// A mapping of artifactStore objects and their corresponding AWS Regions.
	//
	// There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#artifact_stores CodepipelinePipeline#artifact_stores}
	ArtifactStores interface{} `field:"optional" json:"artifactStores" yaml:"artifactStores"`
	// Represents the input of a DisableStageTransition action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#disable_inbound_stage_transitions CodepipelinePipeline#disable_inbound_stage_transitions}
	DisableInboundStageTransitions interface{} `field:"optional" json:"disableInboundStageTransitions" yaml:"disableInboundStageTransitions"`
	// The method that the pipeline will use to handle multiple executions. The default mode is SUPERSEDED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#execution_mode CodepipelinePipeline#execution_mode}
	ExecutionMode *string `field:"optional" json:"executionMode" yaml:"executionMode"`
	// The name of the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#name CodepipelinePipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// CodePipeline provides the following pipeline types, which differ in characteristics and price, so that you can tailor your pipeline features and cost to the needs of your applications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#pipeline_type CodepipelinePipeline#pipeline_type}
	PipelineType *string `field:"optional" json:"pipelineType" yaml:"pipelineType"`
	// Indicates whether to rerun the CodePipeline pipeline after you update it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#restart_execution_on_update CodepipelinePipeline#restart_execution_on_update}
	RestartExecutionOnUpdate interface{} `field:"optional" json:"restartExecutionOnUpdate" yaml:"restartExecutionOnUpdate"`
	// Specifies the tags applied to the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#tags CodepipelinePipeline#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The trigger configuration specifying a type of event, such as Git tags, that starts the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#triggers CodepipelinePipeline#triggers}
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
	// A list that defines the pipeline variables for a pipeline resource.
	//
	// Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9@\-_]+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codepipeline_pipeline#variables CodepipelinePipeline#variables}
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

