// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarmmuterule


type CloudwatchAlarmMuteRuleRule struct {
	// Schedule for the mute to be active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudwatch_alarm_mute_rule#schedule CloudwatchAlarmMuteRule#schedule}
	Schedule *CloudwatchAlarmMuteRuleRuleSchedule `field:"required" json:"schedule" yaml:"schedule"`
}

