// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel


type IoteventsAlarmModelAlarmRule struct {
	// A rule that compares an input property value to a threshold value with a comparison operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotevents_alarm_model#simple_rule IoteventsAlarmModel#simple_rule}
	SimpleRule *IoteventsAlarmModelAlarmRuleSimpleRule `field:"optional" json:"simpleRule" yaml:"simpleRule"`
}

