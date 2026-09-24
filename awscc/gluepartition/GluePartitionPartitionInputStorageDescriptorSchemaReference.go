// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluepartition


type GluePartitionPartitionInputStorageDescriptorSchemaReference struct {
	// A structure that contains schema identity fields. Either this or the SchemaVersionId has to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_partition#schema_id GluePartition#schema_id}
	SchemaId *GluePartitionPartitionInputStorageDescriptorSchemaReferenceSchemaId `field:"optional" json:"schemaId" yaml:"schemaId"`
	// The unique ID assigned to a version of the schema. Either this or the SchemaId has to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_partition#schema_version_id GluePartition#schema_version_id}
	SchemaVersionId *string `field:"optional" json:"schemaVersionId" yaml:"schemaVersionId"`
	// The version number of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_partition#schema_version_number GluePartition#schema_version_number}
	SchemaVersionNumber *float64 `field:"optional" json:"schemaVersionNumber" yaml:"schemaVersionNumber"`
}

