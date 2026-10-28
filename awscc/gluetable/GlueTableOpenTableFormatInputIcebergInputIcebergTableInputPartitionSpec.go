// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpec struct {
	// The partition fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#fields GlueTable#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The spec ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#spec_id GlueTable#spec_id}
	SpecId *float64 `field:"optional" json:"specId" yaml:"specId"`
}

