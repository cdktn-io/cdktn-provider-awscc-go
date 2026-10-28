// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFields struct {
	// The partition field ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#field_id GlueTable#field_id}
	FieldId *float64 `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The partition field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#name GlueTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The source field ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#source_id GlueTable#source_id}
	SourceId *float64 `field:"optional" json:"sourceId" yaml:"sourceId"`
	// The transform to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#transform GlueTable#transform}
	Transform *string `field:"optional" json:"transform" yaml:"transform"`
}

