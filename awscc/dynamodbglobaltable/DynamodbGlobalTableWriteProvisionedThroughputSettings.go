// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableWriteProvisionedThroughputSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dynamodb_global_table#write_capacity_auto_scaling_settings DynamodbGlobalTable#write_capacity_auto_scaling_settings}.
	WriteCapacityAutoScalingSettings *DynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettings `field:"optional" json:"writeCapacityAutoScalingSettings" yaml:"writeCapacityAutoScalingSettings"`
}

