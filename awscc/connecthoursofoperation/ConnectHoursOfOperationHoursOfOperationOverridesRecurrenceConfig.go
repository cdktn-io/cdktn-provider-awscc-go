// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecthoursofoperation


type ConnectHoursOfOperationHoursOfOperationOverridesRecurrenceConfig struct {
	// Pattern for recurring hours of operation overrides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_hours_of_operation#recurrence_pattern ConnectHoursOfOperation#recurrence_pattern}
	RecurrencePattern *ConnectHoursOfOperationHoursOfOperationOverridesRecurrenceConfigRecurrencePattern `field:"optional" json:"recurrencePattern" yaml:"recurrencePattern"`
}

