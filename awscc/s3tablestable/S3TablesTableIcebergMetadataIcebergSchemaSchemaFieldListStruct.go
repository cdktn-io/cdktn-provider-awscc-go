// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableIcebergMetadataIcebergSchemaSchemaFieldListStruct struct {
	// The unique identifier for the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3tables_table#id S3TablesTable#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// The name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3tables_table#name S3TablesTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A Boolean value that specifies whether values are required for each row in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3tables_table#required S3TablesTable#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The field type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3tables_table#type S3TablesTable#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

