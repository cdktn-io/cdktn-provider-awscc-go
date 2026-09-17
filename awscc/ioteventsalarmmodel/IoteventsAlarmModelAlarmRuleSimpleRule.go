// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel


type IoteventsAlarmModelAlarmRuleSimpleRule struct {
	// The comparison operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#comparison_operator IoteventsAlarmModel#comparison_operator}
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value on the left side of the comparison operator.
	//
	// You can specify an ITE input attribute as an input property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#input_property IoteventsAlarmModel#input_property}
	InputProperty *string `field:"optional" json:"inputProperty" yaml:"inputProperty"`
	// The value on the right side of the comparison operator.
	//
	// You can enter a number or specify an ITE input attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#threshold IoteventsAlarmModel#threshold}
	Threshold *string `field:"optional" json:"threshold" yaml:"threshold"`
}

