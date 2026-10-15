// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFields struct {
	// The sort direction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#direction GlueTable#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// The null ordering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#null_order GlueTable#null_order}
	NullOrder *string `field:"optional" json:"nullOrder" yaml:"nullOrder"`
	// The source field ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#source_id GlueTable#source_id}
	SourceId *float64 `field:"optional" json:"sourceId" yaml:"sourceId"`
	// The transform to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#transform GlueTable#transform}
	Transform *string `field:"optional" json:"transform" yaml:"transform"`
}

