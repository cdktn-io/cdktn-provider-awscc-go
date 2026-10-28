// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmActionConfiguration struct {
	// Whether the alarm's tags are included in its action notifications and state-change events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudwatch_alarm#alarm_tags_propagation CloudwatchAlarm#alarm_tags_propagation}
	AlarmTagsPropagation *CloudwatchAlarmActionConfigurationAlarmTagsPropagation `field:"optional" json:"alarmTagsPropagation" yaml:"alarmTagsPropagation"`
}

