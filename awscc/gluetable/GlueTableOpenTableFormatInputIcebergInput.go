// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableOpenTableFormatInputIcebergInput struct {
	// Specifies the Iceberg table configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#iceberg_table_input GlueTable#iceberg_table_input}
	IcebergTableInput *GlueTableOpenTableFormatInputIcebergInputIcebergTableInput `field:"optional" json:"icebergTableInput" yaml:"icebergTableInput"`
	// A required metadata operation. Can only be set to CREATE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#metadata_operation GlueTable#metadata_operation}
	MetadataOperation *string `field:"optional" json:"metadataOperation" yaml:"metadataOperation"`
	// The table version of the Iceberg table. Defaults to 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#version GlueTable#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

