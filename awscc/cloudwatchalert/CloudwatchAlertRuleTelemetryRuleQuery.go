// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert


type CloudwatchAlertRuleTelemetryRuleQuery struct {
	// The query expression. For PromQL, the series value is returned in a field named value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#expression CloudwatchAlert#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The query language. Valid values: SQL, PROMQL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#language CloudwatchAlert#language}
	Language *string `field:"required" json:"language" yaml:"language"`
}

