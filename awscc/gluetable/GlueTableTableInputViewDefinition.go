// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableTableInputViewDefinition struct {
	// The definer of the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#definer GlueTable#definer}
	Definer *string `field:"optional" json:"definer" yaml:"definer"`
	// Indicates whether the view is protected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#is_protected GlueTable#is_protected}
	IsProtected interface{} `field:"optional" json:"isProtected" yaml:"isProtected"`
	// The representations of the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#representations GlueTable#representations}
	Representations interface{} `field:"optional" json:"representations" yaml:"representations"`
	// The sub-objects of the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#sub_objects GlueTable#sub_objects}
	SubObjects *[]*string `field:"optional" json:"subObjects" yaml:"subObjects"`
}

