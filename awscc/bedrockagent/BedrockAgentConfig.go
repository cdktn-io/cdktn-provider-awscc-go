// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockAgentConfig struct {
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
	// Name for a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#agent_name BedrockAgent#agent_name}
	AgentName *string `field:"required" json:"agentName" yaml:"agentName"`
	// List of ActionGroups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#action_groups BedrockAgent#action_groups}
	ActionGroups interface{} `field:"optional" json:"actionGroups" yaml:"actionGroups"`
	// Agent collaboration state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#agent_collaboration BedrockAgent#agent_collaboration}
	AgentCollaboration *string `field:"optional" json:"agentCollaboration" yaml:"agentCollaboration"`
	// List of Agent Collaborators.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#agent_collaborators BedrockAgent#agent_collaborators}
	AgentCollaborators interface{} `field:"optional" json:"agentCollaborators" yaml:"agentCollaborators"`
	// ARN of a IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#agent_resource_role_arn BedrockAgent#agent_resource_role_arn}
	AgentResourceRoleArn *string `field:"optional" json:"agentResourceRoleArn" yaml:"agentResourceRoleArn"`
	// Specifies whether to automatically prepare after creating or updating the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#auto_prepare BedrockAgent#auto_prepare}
	AutoPrepare interface{} `field:"optional" json:"autoPrepare" yaml:"autoPrepare"`
	// A KMS key ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#customer_encryption_key_arn BedrockAgent#customer_encryption_key_arn}
	CustomerEncryptionKeyArn *string `field:"optional" json:"customerEncryptionKeyArn" yaml:"customerEncryptionKeyArn"`
	// Structure for custom orchestration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#custom_orchestration BedrockAgent#custom_orchestration}
	CustomOrchestration *BedrockAgentCustomOrchestration `field:"optional" json:"customOrchestration" yaml:"customOrchestration"`
	// Description of the Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#description BedrockAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// ARN or name of a Bedrock model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#foundation_model BedrockAgent#foundation_model}
	FoundationModel *string `field:"optional" json:"foundationModel" yaml:"foundationModel"`
	// Configuration for a guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#guardrail_configuration BedrockAgent#guardrail_configuration}
	GuardrailConfiguration *BedrockAgentGuardrailConfiguration `field:"optional" json:"guardrailConfiguration" yaml:"guardrailConfiguration"`
	// Max Session Time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#idle_session_ttl_in_seconds BedrockAgent#idle_session_ttl_in_seconds}
	IdleSessionTtlInSeconds *float64 `field:"optional" json:"idleSessionTtlInSeconds" yaml:"idleSessionTtlInSeconds"`
	// Instruction for the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#instruction BedrockAgent#instruction}
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
	// List of Agent Knowledge Bases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#knowledge_bases BedrockAgent#knowledge_bases}
	KnowledgeBases interface{} `field:"optional" json:"knowledgeBases" yaml:"knowledgeBases"`
	// Configuration for memory storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#memory_configuration BedrockAgent#memory_configuration}
	MemoryConfiguration *BedrockAgentMemoryConfiguration `field:"optional" json:"memoryConfiguration" yaml:"memoryConfiguration"`
	// Types of orchestration strategy for agents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#orchestration_type BedrockAgent#orchestration_type}
	OrchestrationType *string `field:"optional" json:"orchestrationType" yaml:"orchestrationType"`
	// Configuration for prompt override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#prompt_override_configuration BedrockAgent#prompt_override_configuration}
	PromptOverrideConfiguration *BedrockAgentPromptOverrideConfiguration `field:"optional" json:"promptOverrideConfiguration" yaml:"promptOverrideConfiguration"`
	// Specifies whether to allow deleting agent while it is in use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#skip_resource_in_use_check_on_delete BedrockAgent#skip_resource_in_use_check_on_delete}
	SkipResourceInUseCheckOnDelete interface{} `field:"optional" json:"skipResourceInUseCheckOnDelete" yaml:"skipResourceInUseCheckOnDelete"`
	// A map of tag keys and values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#tags BedrockAgent#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// A map of tag keys and values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#test_alias_tags BedrockAgent#test_alias_tags}
	TestAliasTags *map[string]*string `field:"optional" json:"testAliasTags" yaml:"testAliasTags"`
}

