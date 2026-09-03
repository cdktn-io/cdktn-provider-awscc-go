// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecthoursofoperation


type ConnectHoursOfOperationHoursOfOperationOverridesRecurrenceConfigRecurrencePattern struct {
	// List of months (1-12) for recurrence pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_hours_of_operation#by_month ConnectHoursOfOperation#by_month}
	ByMonth *[]*float64 `field:"optional" json:"byMonth" yaml:"byMonth"`
	// List of month days (-1 to 31) for recurrence pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_hours_of_operation#by_month_day ConnectHoursOfOperation#by_month_day}
	ByMonthDay *[]*float64 `field:"optional" json:"byMonthDay" yaml:"byMonthDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_hours_of_operation#by_weekday_occurrence ConnectHoursOfOperation#by_weekday_occurrence}.
	ByWeekdayOccurrence *[]*float64 `field:"optional" json:"byWeekdayOccurrence" yaml:"byWeekdayOccurrence"`
	// The frequency of recurrence for hours of operation overrides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_hours_of_operation#frequency ConnectHoursOfOperation#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_hours_of_operation#interval ConnectHoursOfOperation#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

