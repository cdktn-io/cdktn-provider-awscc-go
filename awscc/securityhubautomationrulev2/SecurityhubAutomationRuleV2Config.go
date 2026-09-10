// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrulev2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityhubAutomationRuleV2Config struct {
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
	// A list of actions to be performed when the rule criteria is met.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#actions SecurityhubAutomationRuleV2#actions}
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// Defines the parameters and conditions used to evaluate and filter security findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#criteria SecurityhubAutomationRuleV2#criteria}
	Criteria *SecurityhubAutomationRuleV2Criteria `field:"required" json:"criteria" yaml:"criteria"`
	// A description of the automation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#description SecurityhubAutomationRuleV2#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the automation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#rule_name SecurityhubAutomationRuleV2#rule_name}
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// The value for the rule priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#rule_order SecurityhubAutomationRuleV2#rule_order}
	RuleOrder *float64 `field:"required" json:"ruleOrder" yaml:"ruleOrder"`
	// The status of the automation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#rule_status SecurityhubAutomationRuleV2#rule_status}
	RuleStatus *string `field:"optional" json:"ruleStatus" yaml:"ruleStatus"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_automation_rule_v2#tags SecurityhubAutomationRuleV2#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

