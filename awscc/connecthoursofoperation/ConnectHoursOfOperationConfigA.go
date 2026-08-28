// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecthoursofoperation


type ConnectHoursOfOperationConfigA struct {
	// The day that the hours of operation applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_hours_of_operation#day ConnectHoursOfOperation#day}
	Day *string `field:"required" json:"day" yaml:"day"`
	// The end time that your contact center closes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_hours_of_operation#end_time ConnectHoursOfOperation#end_time}
	EndTime *ConnectHoursOfOperationConfigEndTime `field:"required" json:"endTime" yaml:"endTime"`
	// The start time that your contact center opens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_hours_of_operation#start_time ConnectHoursOfOperation#start_time}
	StartTime *ConnectHoursOfOperationConfigStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

