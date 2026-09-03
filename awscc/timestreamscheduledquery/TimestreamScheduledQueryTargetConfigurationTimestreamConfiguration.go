// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamscheduledquery


type TimestreamScheduledQueryTargetConfigurationTimestreamConfiguration struct {
	// Name of Timestream database to which the query result will be written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_scheduled_query#database_name TimestreamScheduledQuery#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// This is to allow mapping column(s) from the query result to the dimension in the destination table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_scheduled_query#dimension_mappings TimestreamScheduledQuery#dimension_mappings}
	DimensionMappings interface{} `field:"optional" json:"dimensionMappings" yaml:"dimensionMappings"`
	// Name of the measure name column from the query result.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_scheduled_query#measure_name_column TimestreamScheduledQuery#measure_name_column}
	MeasureNameColumn *string `field:"optional" json:"measureNameColumn" yaml:"measureNameColumn"`
	// Specifies how to map measures to multi-measure records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_scheduled_query#mixed_measure_mappings TimestreamScheduledQuery#mixed_measure_mappings}
	MixedMeasureMappings interface{} `field:"optional" json:"mixedMeasureMappings" yaml:"mixedMeasureMappings"`
	// Only one of MixedMeasureMappings or MultiMeasureMappings is to be provided.
	//
	// MultiMeasureMappings can be used to ingest data as multi measures in the derived table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_scheduled_query#multi_measure_mappings TimestreamScheduledQuery#multi_measure_mappings}
	MultiMeasureMappings *TimestreamScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappings `field:"optional" json:"multiMeasureMappings" yaml:"multiMeasureMappings"`
	// Name of Timestream table that the query result will be written to.
	//
	// The table should be within the same database that is provided in Timestream configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_scheduled_query#table_name TimestreamScheduledQuery#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Column from query result that should be used as the time column in destination table.
	//
	// Column type for this should be TIMESTAMP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_scheduled_query#time_column TimestreamScheduledQuery#time_column}
	TimeColumn *string `field:"optional" json:"timeColumn" yaml:"timeColumn"`
}

