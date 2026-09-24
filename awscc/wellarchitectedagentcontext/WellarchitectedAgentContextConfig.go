// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedagentcontext

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WellarchitectedAgentContextConfig struct {
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
	// The free-form content of the Agent Context, supplied as an arbitrary JSON object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_context#content WellarchitectedAgentContext#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The type of the Agent Context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_context#context_type WellarchitectedAgentContext#context_type}
	ContextType *string `field:"required" json:"contextType" yaml:"contextType"`
	// The Amazon Resource Name (ARN) of the parent Agent Profile that owns this context.
	//
	// Pass `!Ref` of the parent AWS::WellArchitected::AgentProfile to flow its ARN here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_context#profile_arn WellarchitectedAgentContext#profile_arn}
	ProfileArn *string `field:"required" json:"profileArn" yaml:"profileArn"`
	// The title of the Agent Context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_context#title WellarchitectedAgentContext#title}
	Title *string `field:"required" json:"title" yaml:"title"`
}

