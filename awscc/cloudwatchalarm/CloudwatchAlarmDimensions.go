// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmDimensions struct {
	// The name of the dimension, from 1–255 characters in length.
	//
	// This dimension name must have been included when the metric was published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#name CloudwatchAlarm#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value for the dimension, from 1–255 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#value CloudwatchAlarm#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

