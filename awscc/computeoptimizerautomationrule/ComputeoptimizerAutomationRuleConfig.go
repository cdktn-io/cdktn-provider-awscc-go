// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeoptimizerautomationrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeoptimizerAutomationRuleConfig struct {
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
	// The name of the automation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#name ComputeoptimizerAutomationRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The types of recommended actions this rule will implement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#recommended_action_types ComputeoptimizerAutomationRule#recommended_action_types}
	RecommendedActionTypes *[]*string `field:"required" json:"recommendedActionTypes" yaml:"recommendedActionTypes"`
	// The type of automation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#rule_type ComputeoptimizerAutomationRule#rule_type}
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	// The schedule configuration for when the rule runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#schedule ComputeoptimizerAutomationRule#schedule}
	Schedule *ComputeoptimizerAutomationRuleSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// The status of the automation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#status ComputeoptimizerAutomationRule#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// Filter criteria that specify which recommended actions qualify for implementation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#criteria ComputeoptimizerAutomationRule#criteria}
	Criteria *ComputeoptimizerAutomationRuleCriteria `field:"optional" json:"criteria" yaml:"criteria"`
	// The description of the automation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#description ComputeoptimizerAutomationRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Organization configuration for organization rules, including rule apply order and account scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#organization_configuration ComputeoptimizerAutomationRule#organization_configuration}
	OrganizationConfiguration *ComputeoptimizerAutomationRuleOrganizationConfiguration `field:"optional" json:"organizationConfiguration" yaml:"organizationConfiguration"`
	// Rule priority within its group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#priority ComputeoptimizerAutomationRule#priority}
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Tags associated with the automation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#tags ComputeoptimizerAutomationRule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

