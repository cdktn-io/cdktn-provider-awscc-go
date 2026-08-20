// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableReplicasPointInTimeRecoverySpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#point_in_time_recovery_enabled DynamodbGlobalTable#point_in_time_recovery_enabled}.
	PointInTimeRecoveryEnabled interface{} `field:"optional" json:"pointInTimeRecoveryEnabled" yaml:"pointInTimeRecoveryEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#recovery_period_in_days DynamodbGlobalTable#recovery_period_in_days}.
	RecoveryPeriodInDays *float64 `field:"optional" json:"recoveryPeriodInDays" yaml:"recoveryPeriodInDays"`
}

