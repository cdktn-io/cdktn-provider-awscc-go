// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamscheduledquery


type TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationDimensionMappings struct {
	// Type for the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_scheduled_query#dimension_value_type TimestreamScheduledQuery#dimension_value_type}
	DimensionValueType *string `field:"optional" json:"dimensionValueType" yaml:"dimensionValueType"`
	// Column name from query result.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_scheduled_query#name TimestreamScheduledQuery#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

