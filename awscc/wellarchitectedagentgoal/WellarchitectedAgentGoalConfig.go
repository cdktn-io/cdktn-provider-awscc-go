// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedagentgoal

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WellarchitectedAgentGoalConfig struct {
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
	// The list of Well-Architected pillars this goal targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_goal#pillars WellarchitectedAgentGoal#pillars}
	Pillars *[]*string `field:"required" json:"pillars" yaml:"pillars"`
	// The Amazon Resource Name (ARN) of the parent Agent Profile that owns this goal.
	//
	// Pass `!Ref` of the parent AWS::WellArchitected::AgentProfile to flow its ARN here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_goal#profile_arn WellarchitectedAgentGoal#profile_arn}
	ProfileArn *string `field:"required" json:"profileArn" yaml:"profileArn"`
	// The title of the Agent Goal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_goal#title WellarchitectedAgentGoal#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// A description of the Agent Goal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_goal#description WellarchitectedAgentGoal#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

