// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instanceeventwindow


type Ec2InstanceEventWindowTimeRanges struct {
	// The hour when the time range ends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_instance_event_window#end_hour Ec2InstanceEventWindow#end_hour}
	EndHour *float64 `field:"optional" json:"endHour" yaml:"endHour"`
	// The day on which the time range ends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_instance_event_window#end_week_day Ec2InstanceEventWindow#end_week_day}
	EndWeekDay *string `field:"optional" json:"endWeekDay" yaml:"endWeekDay"`
	// The hour when the time range begins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_instance_event_window#start_hour Ec2InstanceEventWindow#start_hour}
	StartHour *float64 `field:"optional" json:"startHour" yaml:"startHour"`
	// The day on which the time range begins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_instance_event_window#start_week_day Ec2InstanceEventWindow#start_week_day}
	StartWeekDay *string `field:"optional" json:"startWeekDay" yaml:"startWeekDay"`
}

