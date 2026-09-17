// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationEntityConfigurationsSchema struct {
	// The data type of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#field_data_type GlueConnectionType#field_data_type}
	FieldDataType *string `field:"optional" json:"fieldDataType" yaml:"fieldDataType"`
	// Configuration that defines per-field overrides for filter behavior, allowing individual fields to customize how filter operations are applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#filter_overrides GlueConnectionType#filter_overrides}
	FilterOverrides *GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaFilterOverrides `field:"optional" json:"filterOverrides" yaml:"filterOverrides"`
	// Indicates whether this field can contain null values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#is_nullable GlueConnectionType#is_nullable}
	IsNullable interface{} `field:"optional" json:"isNullable" yaml:"isNullable"`
	// Indicates whether this field can be used for ordering results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#is_orderable GlueConnectionType#is_orderable}
	IsOrderable interface{} `field:"optional" json:"isOrderable" yaml:"isOrderable"`
	// Indicates whether this field can be used for partitioning queries to the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#is_partitionable GlueConnectionType#is_partitionable}
	IsPartitionable interface{} `field:"optional" json:"isPartitionable" yaml:"isPartitionable"`
	// Indicates whether this field can be used in filter predicates when querying data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#is_queryable GlueConnectionType#is_queryable}
	IsQueryable interface{} `field:"optional" json:"isQueryable" yaml:"isQueryable"`
	// The name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#name GlueConnectionType#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The format pattern for parsing date values from API responses. Accepts Java DateTimeFormatter patterns, EPOCH_SECONDS, or EPOCH_MILLIS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#response_date_format GlueConnectionType#response_date_format}
	ResponseDateFormat *string `field:"optional" json:"responseDateFormat" yaml:"responseDateFormat"`
}

