// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableAttributeDefinitions struct {
	// A name for the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dynamodb_table#attribute_name DynamodbTable#attribute_name}
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// The data type for the attribute, where:   +  ``S`` - the attribute is of type String   +  ``N`` - the attribute is of type Number   +  ``B`` - the attribute is of type Binary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dynamodb_table#attribute_type DynamodbTable#attribute_type}
	AttributeType *string `field:"optional" json:"attributeType" yaml:"attributeType"`
}

