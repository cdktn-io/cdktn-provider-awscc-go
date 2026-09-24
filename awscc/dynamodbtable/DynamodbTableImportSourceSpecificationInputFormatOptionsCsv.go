// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableImportSourceSpecificationInputFormatOptionsCsv struct {
	// The delimiter used for separating items in the CSV file being imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_table#delimiter DynamodbTable#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// List of the headers used to specify a common header for all source CSV files being imported.
	//
	// If this field is specified then the first line of each CSV file is treated as data instead of the header. If this field is not specified the the first line of each CSV file is treated as the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_table#header_list DynamodbTable#header_list}
	HeaderList *[]*string `field:"optional" json:"headerList" yaml:"headerList"`
}

