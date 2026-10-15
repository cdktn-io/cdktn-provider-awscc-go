// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFields struct {
	// Documentation for the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#doc GlueTable#doc}
	Doc *string `field:"optional" json:"doc" yaml:"doc"`
	// The field ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#id GlueTable#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// The field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#name GlueTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether the field is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#required GlueTable#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The field type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#type GlueTable#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

