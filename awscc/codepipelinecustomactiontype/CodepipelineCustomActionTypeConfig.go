// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinecustomactiontype

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CodepipelineCustomActionTypeConfig struct {
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
	// The category of the custom action, such as a build action or a test action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_custom_action_type#category CodepipelineCustomActionType#category}
	Category *string `field:"required" json:"category" yaml:"category"`
	// The details of the input artifact for the action, such as its commit ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_custom_action_type#input_artifact_details CodepipelineCustomActionType#input_artifact_details}
	InputArtifactDetails *CodepipelineCustomActionTypeInputArtifactDetails `field:"required" json:"inputArtifactDetails" yaml:"inputArtifactDetails"`
	// The details of the output artifact of the action, such as its commit ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_custom_action_type#output_artifact_details CodepipelineCustomActionType#output_artifact_details}
	OutputArtifactDetails *CodepipelineCustomActionTypeOutputArtifactDetails `field:"required" json:"outputArtifactDetails" yaml:"outputArtifactDetails"`
	// The provider of the service used in the custom action, such as AWS CodeDeploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_custom_action_type#provider_name CodepipelineCustomActionType#provider_name}
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The version identifier of the custom action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_custom_action_type#version CodepipelineCustomActionType#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// The configuration properties for the custom action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_custom_action_type#configuration_properties CodepipelineCustomActionType#configuration_properties}
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// URLs that provide users information about this custom action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_custom_action_type#settings CodepipelineCustomActionType#settings}
	Settings *CodepipelineCustomActionTypeSettings `field:"optional" json:"settings" yaml:"settings"`
	// Any tags assigned to the custom action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_custom_action_type#tags CodepipelineCustomActionType#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

