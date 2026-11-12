// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert


type CloudwatchAlertRule struct {
	// A rule that queries telemetry, evaluates the result against a condition, and sets how missing data is handled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#telemetry_rule CloudwatchAlert#telemetry_rule}
	TelemetryRule *CloudwatchAlertRuleTelemetryRule `field:"required" json:"telemetryRule" yaml:"telemetryRule"`
}

