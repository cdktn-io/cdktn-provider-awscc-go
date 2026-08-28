// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableReplicasReadProvisionedThroughputSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_global_table#read_capacity_auto_scaling_settings DynamodbGlobalTable#read_capacity_auto_scaling_settings}.
	ReadCapacityAutoScalingSettings *DynamodbGlobalTableReplicasReadProvisionedThroughputSettingsReadCapacityAutoScalingSettings `field:"optional" json:"readCapacityAutoScalingSettings" yaml:"readCapacityAutoScalingSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_global_table#read_capacity_units DynamodbGlobalTable#read_capacity_units}.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
}

