// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableTimeToLiveSpecification struct {
	// The name of the TTL attribute used to store the expiration time for items in the table.
	//
	// +  The ``AttributeName`` property is required when enabling the TTL, or when TTL is already enabled.
	//   +  To update this property, you must first disable TTL and then enable TTL with the new attribute name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_table#attribute_name DynamodbTable#attribute_name}
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// Indicates whether TTL is to be enabled (true) or disabled (false) on the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_table#enabled DynamodbTable#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

