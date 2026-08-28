// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkfirewallRuleGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_rule_group#capacity NetworkfirewallRuleGroup#capacity}.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_rule_group#rule_group_name NetworkfirewallRuleGroup#rule_group_name}.
	RuleGroupName *string `field:"required" json:"ruleGroupName" yaml:"ruleGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_rule_group#type NetworkfirewallRuleGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_rule_group#description NetworkfirewallRuleGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_rule_group#rule_group NetworkfirewallRuleGroup#rule_group}.
	RuleGroup *NetworkfirewallRuleGroupRuleGroup `field:"optional" json:"ruleGroup" yaml:"ruleGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_rule_group#summary_configuration NetworkfirewallRuleGroup#summary_configuration}.
	SummaryConfiguration *NetworkfirewallRuleGroupSummaryConfiguration `field:"optional" json:"summaryConfiguration" yaml:"summaryConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_rule_group#tags NetworkfirewallRuleGroup#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

