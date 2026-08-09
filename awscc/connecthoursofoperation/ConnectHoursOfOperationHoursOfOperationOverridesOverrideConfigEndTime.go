// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecthoursofoperation


type ConnectHoursOfOperationHoursOfOperationOverridesOverrideConfigEndTime struct {
	// The hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_hours_of_operation#hours ConnectHoursOfOperation#hours}
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// The minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_hours_of_operation#minutes ConnectHoursOfOperation#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
}

