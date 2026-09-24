// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedagentprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WellarchitectedAgentProfileConfig struct {
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
	// The aggregation configuration entries (account, regions, access role) associated with this profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#aggregation_configuration WellarchitectedAgentProfile#aggregation_configuration}
	AggregationConfiguration interface{} `field:"required" json:"aggregationConfiguration" yaml:"aggregationConfiguration"`
	// The ARN of the IAM role assumed to execute recommendation actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#execution_role_arn WellarchitectedAgentProfile#execution_role_arn}
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of the profile. Unique within the account and used as the last component of the ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#name WellarchitectedAgentProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of Well-Architected pillars to focus on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#pillars WellarchitectedAgentProfile#pillars}
	Pillars *[]*string `field:"required" json:"pillars" yaml:"pillars"`
	// A business overview for the profile used to improve recommendation quality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#business_overview WellarchitectedAgentProfile#business_overview}
	BusinessOverview *string `field:"optional" json:"businessOverview" yaml:"businessOverview"`
	// Whether deletion protection is enabled for the profile.
	//
	// When enabled, the profile cannot be deleted until deletion protection is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#deletion_protection WellarchitectedAgentProfile#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// A description of the profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#description WellarchitectedAgentProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The human-readable display name of the profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#display_name WellarchitectedAgentProfile#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Key-value pairs to associate with the Agent Profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#tags WellarchitectedAgentProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

