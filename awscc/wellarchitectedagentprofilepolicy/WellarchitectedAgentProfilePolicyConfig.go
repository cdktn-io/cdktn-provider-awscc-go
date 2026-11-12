// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedagentprofilepolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WellarchitectedAgentProfilePolicyConfig struct {
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
	// The IAM resource-based policy document as a JSON string. Allow-only statements, no negation condition operators, maximum 51200 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/wellarchitected_agent_profile_policy#policy WellarchitectedAgentProfilePolicy#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// The Amazon Resource Name (ARN) of the parent Agent Profile the policy is attached to.
	//
	// Pass `!GetAtt` (or `!Ref`) of the parent AWS::WellArchitected::AgentProfile to flow its ARN here. Also serves as the primary identifier of this policy attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/wellarchitected_agent_profile_policy#profile_arn WellarchitectedAgentProfilePolicy#profile_arn}
	ProfileArn *string `field:"required" json:"profileArn" yaml:"profileArn"`
}

