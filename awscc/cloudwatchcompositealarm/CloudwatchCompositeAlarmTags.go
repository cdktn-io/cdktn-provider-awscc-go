// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchcompositealarm


type CloudwatchCompositeAlarmTags struct {
	// A unique identifier for the tag.
	//
	// The combination of tag keys and values can help you organize and categorize your resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_composite_alarm#key CloudwatchCompositeAlarm#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the specified tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudwatch_composite_alarm#value CloudwatchCompositeAlarm#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

