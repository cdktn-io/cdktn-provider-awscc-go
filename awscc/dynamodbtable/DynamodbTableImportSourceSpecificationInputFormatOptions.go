// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableImportSourceSpecificationInputFormatOptions struct {
	// The options for imported source files in CSV format. The values are Delimiter and HeaderList.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_table#csv DynamodbTable#csv}
	Csv *DynamodbTableImportSourceSpecificationInputFormatOptionsCsv `field:"optional" json:"csv" yaml:"csv"`
}

