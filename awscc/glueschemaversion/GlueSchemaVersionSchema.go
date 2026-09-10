// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueschemaversion


type GlueSchemaVersionSchema struct {
	// Name of the registry to identify where the Schema is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_schema_version#registry_name GlueSchemaVersion#registry_name}
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// Amazon Resource Name for the Schema. This attribute can be used to uniquely represent the Schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_schema_version#schema_arn GlueSchemaVersion#schema_arn}
	SchemaArn *string `field:"optional" json:"schemaArn" yaml:"schemaArn"`
	// Name of the schema. This parameter requires RegistryName to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_schema_version#schema_name GlueSchemaVersion#schema_name}
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

