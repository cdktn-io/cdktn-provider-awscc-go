// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarmmuterule


type CloudwatchAlarmMuteRuleMuteTargets struct {
	// The alarm names to be mute by the AlarmMuteRule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudwatch_alarm_mute_rule#alarm_names CloudwatchAlarmMuteRule#alarm_names}
	AlarmNames *[]*string `field:"optional" json:"alarmNames" yaml:"alarmNames"`
}

