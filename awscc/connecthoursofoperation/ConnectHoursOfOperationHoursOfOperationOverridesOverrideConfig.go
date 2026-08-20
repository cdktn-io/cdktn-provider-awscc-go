// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecthoursofoperation


type ConnectHoursOfOperationHoursOfOperationOverridesOverrideConfig struct {
	// The day that the hours of operation override applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_hours_of_operation#day ConnectHoursOfOperation#day}
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The new end time that your contact center closes for the overriden days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_hours_of_operation#end_time ConnectHoursOfOperation#end_time}
	EndTime *ConnectHoursOfOperationHoursOfOperationOverridesOverrideConfigEndTime `field:"optional" json:"endTime" yaml:"endTime"`
	// The new start time that your contact center opens for the overriden days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_hours_of_operation#start_time ConnectHoursOfOperation#start_time}
	StartTime *ConnectHoursOfOperationHoursOfOperationOverridesOverrideConfigStartTime `field:"optional" json:"startTime" yaml:"startTime"`
}

