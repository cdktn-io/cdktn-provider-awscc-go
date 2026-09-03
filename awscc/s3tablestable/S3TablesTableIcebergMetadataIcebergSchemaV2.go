// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableIcebergMetadataIcebergSchemaV2 struct {
	// A list of field IDs that are used as the identifier fields for the table.
	//
	// Identifier fields uniquely identify a row in the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3tables_table#identifier_field_ids S3TablesTable#identifier_field_ids}
	IdentifierFieldIds *[]*float64 `field:"optional" json:"identifierFieldIds" yaml:"identifierFieldIds"`
	// An optional unique identifier for the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3tables_table#schema_id S3TablesTable#schema_id}
	SchemaId *float64 `field:"optional" json:"schemaId" yaml:"schemaId"`
	// The schema fields for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3tables_table#schema_v2_field_list S3TablesTable#schema_v2_field_list}
	SchemaV2FieldList interface{} `field:"optional" json:"schemaV2FieldList" yaml:"schemaV2FieldList"`
	// The type of the top-level schema, which is always 'struct'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3tables_table#schema_v2_field_type S3TablesTable#schema_v2_field_type}
	SchemaV2FieldType *string `field:"optional" json:"schemaV2FieldType" yaml:"schemaV2FieldType"`
}

