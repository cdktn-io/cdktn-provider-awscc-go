// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rbinrule


type RbinRuleRetentionPeriod struct {
	// The retention period unit of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rbin_rule#retention_period_unit RbinRule#retention_period_unit}
	RetentionPeriodUnit *string `field:"required" json:"retentionPeriodUnit" yaml:"retentionPeriodUnit"`
	// The retention period value of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rbin_rule#retention_period_value RbinRule#retention_period_value}
	RetentionPeriodValue *float64 `field:"required" json:"retentionPeriodValue" yaml:"retentionPeriodValue"`
}

