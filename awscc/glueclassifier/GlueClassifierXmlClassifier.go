// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueclassifier


type GlueClassifierXmlClassifier struct {
	// An identifier of the data format that the classifier matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_classifier#classification GlueClassifier#classification}
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// The name of the classifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_classifier#name GlueClassifier#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The XML tag designating the element that contains each record in an XML document being parsed.
	//
	// This can't identify a self-closing element (closed by />). An empty row element that contains only attributes can be parsed as long as it ends with a closing tag (for example, <row item_a="A" item_b="B"></row> is okay, but <row item_a="A" item_b="B" /> is not).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_classifier#row_tag GlueClassifier#row_tag}
	RowTag *string `field:"optional" json:"rowTag" yaml:"rowTag"`
}

