// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableTableInputViewDefinitionRepresentations struct {
	// The dialect of the view representation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#dialect GlueTable#dialect}
	Dialect *string `field:"optional" json:"dialect" yaml:"dialect"`
	// The version of the dialect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#dialect_version GlueTable#dialect_version}
	DialectVersion *string `field:"optional" json:"dialectVersion" yaml:"dialectVersion"`
	// The connection used for validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#validation_connection GlueTable#validation_connection}
	ValidationConnection *string `field:"optional" json:"validationConnection" yaml:"validationConnection"`
	// The expanded text of the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#view_expanded_text GlueTable#view_expanded_text}
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	// The original text of the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#view_original_text GlueTable#view_original_text}
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

