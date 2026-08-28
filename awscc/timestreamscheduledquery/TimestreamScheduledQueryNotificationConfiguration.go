// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamscheduledquery


type TimestreamScheduledQueryNotificationConfiguration struct {
	// SNS configuration for notification upon scheduled query execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/timestream_scheduled_query#sns_configuration TimestreamScheduledQuery#sns_configuration}
	SnsConfiguration *TimestreamScheduledQueryNotificationConfigurationSnsConfiguration `field:"required" json:"snsConfiguration" yaml:"snsConfiguration"`
}

