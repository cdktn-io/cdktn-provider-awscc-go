// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableTableInput struct {
	// A description of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#description GlueTable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The table name. For Hive compatibility, this is folded to lowercase when it is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#name GlueTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The table owner. Included for Apache Hive compatibility. Not used in the normal course of AWS Glue operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#owner GlueTable#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// These key-value pairs define properties associated with the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#parameters GlueTable#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// A list of columns by which the table is partitioned. Only primitive types are supported as partition keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#partition_keys GlueTable#partition_keys}
	PartitionKeys interface{} `field:"optional" json:"partitionKeys" yaml:"partitionKeys"`
	// The retention time for this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#retention GlueTable#retention}
	Retention *float64 `field:"optional" json:"retention" yaml:"retention"`
	// A storage descriptor containing information about the physical storage of this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#storage_descriptor GlueTable#storage_descriptor}
	StorageDescriptor *GlueTableTableInputStorageDescriptor `field:"optional" json:"storageDescriptor" yaml:"storageDescriptor"`
	// The type of this table.
	//
	// AWS Glue will create tables with the EXTERNAL_TABLE type. Other services, such as Athena, may create tables with additional table types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#table_type GlueTable#table_type}
	TableType *string `field:"optional" json:"tableType" yaml:"tableType"`
	// A TableIdentifier structure that describes a target table for resource linking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#target_table GlueTable#target_table}
	TargetTable *GlueTableTableInputTargetTable `field:"optional" json:"targetTable" yaml:"targetTable"`
	// The view definition for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#view_definition GlueTable#view_definition}
	ViewDefinition *GlueTableTableInputViewDefinition `field:"optional" json:"viewDefinition" yaml:"viewDefinition"`
	// Included for Apache Hive compatibility. Not used in the normal course of AWS Glue operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#view_expanded_text GlueTable#view_expanded_text}
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	// Included for Apache Hive compatibility. Not used in the normal course of AWS Glue operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#view_original_text GlueTable#view_original_text}
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

