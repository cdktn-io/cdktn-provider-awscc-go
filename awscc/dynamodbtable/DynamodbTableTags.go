// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableTags struct {
	// The key of the tag.
	//
	// Tag keys are case sensitive. Each DynamoDB table can only have up to one tag with the same key. If you try to add an existing tag (same key), the existing tag value will be updated to the new value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dynamodb_table#key DynamodbTable#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag. Tag values are case-sensitive and can be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dynamodb_table#value DynamodbTable#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

