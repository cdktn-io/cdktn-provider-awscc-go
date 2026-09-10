// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casescaserule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CasesCaseRuleConfig struct {
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
	// A descriptive name for the case rule.
	//
	// Must be unique within the domain and should clearly indicate the rule's purpose (e.g., 'Priority Field Required for Urgent Cases').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_case_rule#name CasesCaseRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Defines the rule behavior and conditions.
	//
	// Specifies the rule type and the conditions under which it applies. In the Amazon Connect admin website, this corresponds to case field conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_case_rule#rule CasesCaseRule#rule}
	Rule *CasesCaseRuleRule `field:"required" json:"rule" yaml:"rule"`
	// A description explaining the purpose and behavior of this case rule.
	//
	// Helps administrators understand when and why this rule applies to case fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_case_rule#description CasesCaseRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the Cases domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_case_rule#domain_id CasesCaseRule#domain_id}
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// The tags that you attach to this case rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_case_rule#tags CasesCaseRule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

