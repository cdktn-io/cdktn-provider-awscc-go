// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableIcebergMetadataIcebergSchema struct {
	// Contains details about the schema for an Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3tables_table#schema_field_list S3TablesTable#schema_field_list}
	SchemaFieldList interface{} `field:"optional" json:"schemaFieldList" yaml:"schemaFieldList"`
}

