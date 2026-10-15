// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableOpenTableFormatInputIcebergInputIcebergTableInputSchema struct {
	// The fields in the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#fields GlueTable#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The identifier field IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#identifier_field_ids GlueTable#identifier_field_ids}
	IdentifierFieldIds *[]*float64 `field:"optional" json:"identifierFieldIds" yaml:"identifierFieldIds"`
	// The schema ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#schema_id GlueTable#schema_id}
	SchemaId *float64 `field:"optional" json:"schemaId" yaml:"schemaId"`
	// The type of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#type GlueTable#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

