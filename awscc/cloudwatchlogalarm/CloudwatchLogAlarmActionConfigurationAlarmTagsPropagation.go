// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogalarm


type CloudwatchLogAlarmActionConfigurationAlarmTagsPropagation struct {
	// Whether the alarm's tags are included in its action notifications and state-change events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudwatch_log_alarm#enabled CloudwatchLogAlarm#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The tag keys to leave out. Cannot be specified together with IncludedTagKeys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudwatch_log_alarm#excluded_tag_keys CloudwatchLogAlarm#excluded_tag_keys}
	ExcludedTagKeys *[]*string `field:"optional" json:"excludedTagKeys" yaml:"excludedTagKeys"`
	// The only tag keys to include. Cannot be specified together with ExcludedTagKeys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudwatch_log_alarm#included_tag_keys CloudwatchLogAlarm#included_tag_keys}
	IncludedTagKeys *[]*string `field:"optional" json:"includedTagKeys" yaml:"includedTagKeys"`
}

