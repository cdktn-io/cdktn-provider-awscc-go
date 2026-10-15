// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableOpenTableFormatInputIcebergInputIcebergTableInput struct {
	// The location of the Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#location GlueTable#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The partition spec for the Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#partition_spec GlueTable#partition_spec}
	PartitionSpec *GlueTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpec `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
	// The properties for the Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#properties GlueTable#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The schema for the Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#schema GlueTable#schema}
	Schema *GlueTableOpenTableFormatInputIcebergInputIcebergTableInputSchema `field:"optional" json:"schema" yaml:"schema"`
	// The write order for the Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#write_order GlueTable#write_order}
	WriteOrder *GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrder `field:"optional" json:"writeOrder" yaml:"writeOrder"`
}

