// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableOpenTableFormatInput struct {
	// Specifies an IcebergInput structure that defines an Apache Iceberg metadata table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#iceberg_input GlueTable#iceberg_input}
	IcebergInput *GlueTableOpenTableFormatInputIcebergInput `field:"optional" json:"icebergInput" yaml:"icebergInput"`
}

