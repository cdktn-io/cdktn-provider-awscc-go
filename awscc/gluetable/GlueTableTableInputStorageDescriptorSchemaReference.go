// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableTableInputStorageDescriptorSchemaReference struct {
	// A structure that contains schema identity fields. Either this or the SchemaVersionId has to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#schema_id GlueTable#schema_id}
	SchemaId *GlueTableTableInputStorageDescriptorSchemaReferenceSchemaId `field:"optional" json:"schemaId" yaml:"schemaId"`
	// The unique ID assigned to a version of the schema. Either this or the SchemaId has to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#schema_version_id GlueTable#schema_version_id}
	SchemaVersionId *string `field:"optional" json:"schemaVersionId" yaml:"schemaVersionId"`
	// The version number of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#schema_version_number GlueTable#schema_version_number}
	SchemaVersionNumber *float64 `field:"optional" json:"schemaVersionNumber" yaml:"schemaVersionNumber"`
}

