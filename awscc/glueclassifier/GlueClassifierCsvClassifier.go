// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueclassifier


type GlueClassifierCsvClassifier struct {
	// Enables the processing of files that contain only one column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_classifier#allow_single_column GlueClassifier#allow_single_column}
	AllowSingleColumn interface{} `field:"optional" json:"allowSingleColumn" yaml:"allowSingleColumn"`
	// Indicates whether the CSV file contains custom data types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_classifier#contains_custom_datatype GlueClassifier#contains_custom_datatype}
	ContainsCustomDatatype *[]*string `field:"optional" json:"containsCustomDatatype" yaml:"containsCustomDatatype"`
	// Indicates whether the CSV file contains a header.
	//
	// A value of UNKNOWN specifies that the classifier will detect whether the CSV file contains headings. A value of PRESENT specifies that the CSV file contains headings. A value of ABSENT specifies that the CSV file does not contain headings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_classifier#contains_header GlueClassifier#contains_header}
	ContainsHeader *string `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// Enables the configuration of custom data types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_classifier#custom_datatype_configured GlueClassifier#custom_datatype_configured}
	CustomDatatypeConfigured interface{} `field:"optional" json:"customDatatypeConfigured" yaml:"customDatatypeConfigured"`
	// A custom symbol to denote what separates each column entry in the row.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_classifier#delimiter GlueClassifier#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Specifies not to trim values before identifying the type of column values. The default value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_classifier#disable_value_trimming GlueClassifier#disable_value_trimming}
	DisableValueTrimming interface{} `field:"optional" json:"disableValueTrimming" yaml:"disableValueTrimming"`
	// A list of strings representing column names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_classifier#header GlueClassifier#header}
	Header *[]*string `field:"optional" json:"header" yaml:"header"`
	// The name of the classifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_classifier#name GlueClassifier#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A custom symbol to denote what combines content into a single column value.
	//
	// It must be different from the column delimiter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_classifier#quote_symbol GlueClassifier#quote_symbol}
	QuoteSymbol *string `field:"optional" json:"quoteSymbol" yaml:"quoteSymbol"`
}

