// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluepartition


type GluePartitionPartitionInputStorageDescriptorSchemaReferenceSchemaId struct {
	// The name of the schema registry that contains the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_partition#registry_name GluePartition#registry_name}
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// The Amazon Resource Name (ARN) of the schema. One of SchemaArn or SchemaName has to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_partition#schema_arn GluePartition#schema_arn}
	SchemaArn *string `field:"optional" json:"schemaArn" yaml:"schemaArn"`
	// The name of the schema. One of SchemaArn or SchemaName has to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_partition#schema_name GluePartition#schema_name}
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

