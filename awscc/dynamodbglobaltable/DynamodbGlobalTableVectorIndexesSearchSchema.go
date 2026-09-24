// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableVectorIndexesSearchSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#attribute_name DynamodbGlobalTable#attribute_name}.
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#search_schema_element_type DynamodbGlobalTable#search_schema_element_type}.
	SearchSchemaElementType *string `field:"optional" json:"searchSchemaElementType" yaml:"searchSchemaElementType"`
}

