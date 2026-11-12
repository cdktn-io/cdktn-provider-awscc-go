// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedagentorganizationalprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WellarchitectedAgentOrganizationalProfileConfig struct {
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
	// The name of the organizational profile. Unique within the account and used as the last component of the ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/wellarchitected_agent_organizational_profile#name WellarchitectedAgentOrganizationalProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of Well-Architected pillars to focus on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/wellarchitected_agent_organizational_profile#pillars WellarchitectedAgentOrganizationalProfile#pillars}
	Pillars *[]*string `field:"required" json:"pillars" yaml:"pillars"`
	// A business overview for the organizational profile used to improve recommendation quality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/wellarchitected_agent_organizational_profile#business_overview WellarchitectedAgentOrganizationalProfile#business_overview}
	BusinessOverview *string `field:"optional" json:"businessOverview" yaml:"businessOverview"`
	// Whether deletion protection is enabled for the organizational profile.
	//
	// When enabled, the profile cannot be deleted until deletion protection is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/wellarchitected_agent_organizational_profile#deletion_protection WellarchitectedAgentOrganizationalProfile#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// A description of the organizational profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/wellarchitected_agent_organizational_profile#description WellarchitectedAgentOrganizationalProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The human-readable display name of the organizational profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/wellarchitected_agent_organizational_profile#display_name WellarchitectedAgentOrganizationalProfile#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Key-value pairs to associate with the Agent Organizational Profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/wellarchitected_agent_organizational_profile#tags WellarchitectedAgentOrganizationalProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

