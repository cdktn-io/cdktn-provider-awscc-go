// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchinsightrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchInsightRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_insight_rule#rule_body CloudwatchInsightRule#rule_body}.
	RuleBody *string `field:"required" json:"ruleBody" yaml:"ruleBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_insight_rule#rule_name CloudwatchInsightRule#rule_name}.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_insight_rule#rule_state CloudwatchInsightRule#rule_state}.
	RuleState *string `field:"required" json:"ruleState" yaml:"ruleState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_insight_rule#apply_on_transformed_logs CloudwatchInsightRule#apply_on_transformed_logs}.
	ApplyOnTransformedLogs interface{} `field:"optional" json:"applyOnTransformedLogs" yaml:"applyOnTransformedLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_insight_rule#tags CloudwatchInsightRule#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

