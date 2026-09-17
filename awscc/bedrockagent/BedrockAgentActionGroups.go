// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentActionGroups struct {
	// Type of Executors for an Action Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#action_group_executor BedrockAgent#action_group_executor}
	ActionGroupExecutor *BedrockAgentActionGroupsActionGroupExecutor `field:"optional" json:"actionGroupExecutor" yaml:"actionGroupExecutor"`
	// Name of the action group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#action_group_name BedrockAgent#action_group_name}
	ActionGroupName *string `field:"optional" json:"actionGroupName" yaml:"actionGroupName"`
	// State of the action group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#action_group_state BedrockAgent#action_group_state}
	ActionGroupState *string `field:"optional" json:"actionGroupState" yaml:"actionGroupState"`
	// Contains information about the API Schema for the Action Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#api_schema BedrockAgent#api_schema}
	ApiSchema *BedrockAgentActionGroupsApiSchema `field:"optional" json:"apiSchema" yaml:"apiSchema"`
	// Description of action group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#description BedrockAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Schema of Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#function_schema BedrockAgent#function_schema}
	FunctionSchema *BedrockAgentActionGroupsFunctionSchema `field:"optional" json:"functionSchema" yaml:"functionSchema"`
	// Action Group Signature for a BuiltIn Action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#parent_action_group_signature BedrockAgent#parent_action_group_signature}
	ParentActionGroupSignature *string `field:"optional" json:"parentActionGroupSignature" yaml:"parentActionGroupSignature"`
	// Specifies whether to allow deleting action group while it is in use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_agent#skip_resource_in_use_check_on_delete BedrockAgent#skip_resource_in_use_check_on_delete}
	SkipResourceInUseCheckOnDelete interface{} `field:"optional" json:"skipResourceInUseCheckOnDelete" yaml:"skipResourceInUseCheckOnDelete"`
}

