// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableVectorIndexesSearchSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_table#attribute_name DynamodbTable#attribute_name}.
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_table#search_schema_element_type DynamodbTable#search_schema_element_type}.
	SearchSchemaElementType *string `field:"optional" json:"searchSchemaElementType" yaml:"searchSchemaElementType"`
}

