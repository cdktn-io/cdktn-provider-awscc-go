// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamscheduledquery


type TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappings struct {
	// Refers to the value of the measure name in a result row.
	//
	// This field is required if MeasureNameColumn is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_scheduled_query#measure_name TimestreamScheduledQuery#measure_name}
	MeasureName *string `field:"optional" json:"measureName" yaml:"measureName"`
	// Type of the value that is to be read from SourceColumn. If the mapping is for MULTI, use MeasureValueType.MULTI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_scheduled_query#measure_value_type TimestreamScheduledQuery#measure_value_type}
	MeasureValueType *string `field:"optional" json:"measureValueType" yaml:"measureValueType"`
	// Required. Attribute mappings to be used for mapping query results to ingest data for multi-measure attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_scheduled_query#multi_measure_attribute_mappings TimestreamScheduledQuery#multi_measure_attribute_mappings}
	MultiMeasureAttributeMappings interface{} `field:"optional" json:"multiMeasureAttributeMappings" yaml:"multiMeasureAttributeMappings"`
	// This field refers to the source column from which the measure value is to be read for result materialization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_scheduled_query#source_column TimestreamScheduledQuery#source_column}
	SourceColumn *string `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// Target measure name to be used.
	//
	// If not provided, the target measure name by default would be MeasureName if provided, or SourceColumn otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_scheduled_query#target_measure_name TimestreamScheduledQuery#target_measure_name}
	TargetMeasureName *string `field:"optional" json:"targetMeasureName" yaml:"targetMeasureName"`
}

