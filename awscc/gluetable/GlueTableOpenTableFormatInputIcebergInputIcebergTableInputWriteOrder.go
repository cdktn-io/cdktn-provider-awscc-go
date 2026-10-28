// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrder struct {
	// The sort fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#fields GlueTable#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The order ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#order_id GlueTable#order_id}
	OrderId *float64 `field:"optional" json:"orderId" yaml:"orderId"`
}

