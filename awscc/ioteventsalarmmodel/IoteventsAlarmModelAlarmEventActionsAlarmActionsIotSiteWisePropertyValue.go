// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWisePropertyValue struct {
	// The quality of the asset property value. The value must be ``'GOOD'``, ``'BAD'``, or ``'UNCERTAIN'``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_alarm_model#quality IoteventsAlarmModel#quality}
	Quality *string `field:"optional" json:"quality" yaml:"quality"`
	// The timestamp associated with the asset property value. The default is the current event time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_alarm_model#timestamp IoteventsAlarmModel#timestamp}
	Timestamp *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWisePropertyValueTimestamp `field:"optional" json:"timestamp" yaml:"timestamp"`
	// The value to send to an asset property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_alarm_model#value IoteventsAlarmModel#value}
	Value *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWisePropertyValueValue `field:"optional" json:"value" yaml:"value"`
}

