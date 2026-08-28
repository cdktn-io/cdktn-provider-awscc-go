// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTablePointInTimeRecoverySpecification struct {
	// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#point_in_time_recovery_enabled DynamodbTable#point_in_time_recovery_enabled}
	PointInTimeRecoveryEnabled interface{} `field:"optional" json:"pointInTimeRecoveryEnabled" yaml:"pointInTimeRecoveryEnabled"`
	// The number of preceding days for which continuous backups are taken and maintained.
	//
	// Your table data is only recoverable to any point-in-time from within the configured recovery period. This parameter is optional. If no value is provided, the value will default to 35.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#recovery_period_in_days DynamodbTable#recovery_period_in_days}
	RecoveryPeriodInDays *float64 `field:"optional" json:"recoveryPeriodInDays" yaml:"recoveryPeriodInDays"`
}

