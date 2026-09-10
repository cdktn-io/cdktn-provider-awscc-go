// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinecustomactiontype


type CodepipelineCustomActionTypeConfigurationProperties struct {
	// The description of the action configuration property that is displayed to users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_custom_action_type#description CodepipelineCustomActionType#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the configuration property is a key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_custom_action_type#key CodepipelineCustomActionType#key}
	Key interface{} `field:"optional" json:"key" yaml:"key"`
	// The name of the action configuration property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_custom_action_type#name CodepipelineCustomActionType#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates that the property is used with PollForJobs.
	//
	// When creating a custom action, an action can have up to one queryable property. If it has one, that property must be both required and not secret.If you create a pipeline with a custom action type, and that custom action contains a queryable property, the value for that configuration property is subject to other restrictions. The value must be less than or equal to twenty (20) characters. The value can contain only alphanumeric characters, underscores, and hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_custom_action_type#queryable CodepipelineCustomActionType#queryable}
	Queryable interface{} `field:"optional" json:"queryable" yaml:"queryable"`
	// Whether the configuration property is a required value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_custom_action_type#required CodepipelineCustomActionType#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Whether the configuration property is secret. Secrets are hidden from all calls except for GetJobDetails, GetThirdPartyJobDetails, PollForJobs, and PollForThirdPartyJobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_custom_action_type#secret CodepipelineCustomActionType#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// The type of the configuration property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_custom_action_type#type CodepipelineCustomActionType#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

