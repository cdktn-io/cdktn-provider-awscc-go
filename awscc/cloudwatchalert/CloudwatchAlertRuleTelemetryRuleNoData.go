// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert


type CloudwatchAlertRuleTelemetryRuleNoData struct {
	// The state to report when an evaluation returns no data. Valid values: OK, WARNING, CRITICAL, NODATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#treat_as CloudwatchAlert#treat_as}
	TreatAs *string `field:"optional" json:"treatAs" yaml:"treatAs"`
}

