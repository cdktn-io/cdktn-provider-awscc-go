// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IoteventsAlarmModelConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Defines when your alarm is invoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#alarm_rule IoteventsAlarmModel#alarm_rule}
	AlarmRule *IoteventsAlarmModelAlarmRule `field:"required" json:"alarmRule" yaml:"alarmRule"`
	// The ARN of the IAM role that allows the alarm to perform actions and access AWS resources.
	//
	// For more information, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#role_arn IoteventsAlarmModel#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Contains the configuration information of alarm state changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#alarm_capabilities IoteventsAlarmModel#alarm_capabilities}
	AlarmCapabilities *IoteventsAlarmModelAlarmCapabilities `field:"optional" json:"alarmCapabilities" yaml:"alarmCapabilities"`
	// Contains information about one or more alarm actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#alarm_event_actions IoteventsAlarmModel#alarm_event_actions}
	AlarmEventActions *IoteventsAlarmModelAlarmEventActions `field:"optional" json:"alarmEventActions" yaml:"alarmEventActions"`
	// The description of the alarm model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#alarm_model_description IoteventsAlarmModel#alarm_model_description}
	AlarmModelDescription *string `field:"optional" json:"alarmModelDescription" yaml:"alarmModelDescription"`
	// The name of the alarm model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#alarm_model_name IoteventsAlarmModel#alarm_model_name}
	AlarmModelName *string `field:"optional" json:"alarmModelName" yaml:"alarmModelName"`
	// An input attribute used as a key to create an alarm.
	//
	// ITE routes [inputs](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Input.html) associated with this key to the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#key IoteventsAlarmModel#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A non-negative integer that reflects the severity level of the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#severity IoteventsAlarmModel#severity}
	Severity *float64 `field:"optional" json:"severity" yaml:"severity"`
	// A list of key-value pairs that contain metadata for the alarm model.
	//
	// The tags help you manage the alarm model. For more information, see [Tagging your resources](https://docs.aws.amazon.com/iotevents/latest/developerguide/tagging-iotevents.html) in the *Developer Guide*.
	//  You can create up to 50 tags for one alarm model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#tags IoteventsAlarmModel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

