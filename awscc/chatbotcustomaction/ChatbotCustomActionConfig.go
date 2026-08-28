// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chatbotcustomaction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChatbotCustomActionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/chatbot_custom_action#action_name ChatbotCustomAction#action_name}.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/chatbot_custom_action#definition ChatbotCustomAction#definition}.
	Definition *ChatbotCustomActionDefinition `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/chatbot_custom_action#alias_name ChatbotCustomAction#alias_name}.
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/chatbot_custom_action#attachments ChatbotCustomAction#attachments}.
	Attachments interface{} `field:"optional" json:"attachments" yaml:"attachments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/chatbot_custom_action#tags ChatbotCustomAction#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

