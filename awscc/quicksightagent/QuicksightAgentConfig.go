// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightagent

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightAgentConfig struct {
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
	// The unique identifier for the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#agent_id QuicksightAgent#agent_id}
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
	// The ID of the Amazon Web Services account where the agent is being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#aws_account_id QuicksightAgent#aws_account_id}
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The display name of the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#name QuicksightAgent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of ActionConnector ARNs (max 10) attached to the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#action_connectors QuicksightAgent#action_connectors}
	ActionConnectors *[]*string `field:"optional" json:"actionConnectors" yaml:"actionConnectors"`
	// The lifecycle stage of the agent. PREVIEW or PUBLISHED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#agent_lifecycle QuicksightAgent#agent_lifecycle}
	AgentLifecycle *string `field:"optional" json:"agentLifecycle" yaml:"agentLifecycle"`
	// Custom prompt configuration. Specify either ExistingPrompt or NewPrompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#custom_prompt_input QuicksightAgent#custom_prompt_input}
	CustomPromptInput *QuicksightAgentCustomPromptInput `field:"optional" json:"customPromptInput" yaml:"customPromptInput"`
	// A description of the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#description QuicksightAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The icon identifier for the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#icon_id QuicksightAgent#icon_id}
	IconId *string `field:"optional" json:"iconId" yaml:"iconId"`
	// A list of Space ARNs (max 10) attached to the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#spaces QuicksightAgent#spaces}
	Spaces *[]*string `field:"optional" json:"spaces" yaml:"spaces"`
	// A list of up to 3 starter prompts displayed to users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#starter_prompts QuicksightAgent#starter_prompts}
	StarterPrompts *[]*string `field:"optional" json:"starterPrompts" yaml:"starterPrompts"`
	// A list of key-value pairs to associate with the agent resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#tags QuicksightAgent#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The welcome message displayed when a user opens the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_agent#welcome_message QuicksightAgent#welcome_message}
	WelcomeMessage *string `field:"optional" json:"welcomeMessage" yaml:"welcomeMessage"`
}

