// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel


type IoteventsAlarmModelAlarmCapabilitiesAcknowledgeFlow struct {
	// The value must be ``TRUE`` or ``FALSE``.
	//
	// If ``TRUE``, you receive a notification when the alarm state changes. You must choose to acknowledge the notification before the alarm state can return to ``NORMAL``. If ``FALSE``, you won't receive notifications. The alarm automatically changes to the ``NORMAL`` state when the input property value returns to the specified range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#enabled IoteventsAlarmModel#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

