// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableGlobalSecondaryIndexesWarmThroughput struct {
	// Represents the number of read operations your base table can instantaneously support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#read_units_per_second DynamodbTable#read_units_per_second}
	ReadUnitsPerSecond *float64 `field:"optional" json:"readUnitsPerSecond" yaml:"readUnitsPerSecond"`
	// Represents the number of write operations your base table can instantaneously support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#write_units_per_second DynamodbTable#write_units_per_second}
	WriteUnitsPerSecond *float64 `field:"optional" json:"writeUnitsPerSecond" yaml:"writeUnitsPerSecond"`
}

