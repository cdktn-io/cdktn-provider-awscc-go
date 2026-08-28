// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableGlobalSecondaryIndexesOnDemandThroughput struct {
	// Maximum number of read request units for the specified table.
	//
	// To specify a maximum ``OnDemandThroughput`` on your table, set the value of ``MaxReadRequestUnits`` as greater than or equal to 1. To remove the maximum ``OnDemandThroughput`` that is currently set on your table, set the value of ``MaxReadRequestUnits`` to -1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#max_read_request_units DynamodbTable#max_read_request_units}
	MaxReadRequestUnits *float64 `field:"optional" json:"maxReadRequestUnits" yaml:"maxReadRequestUnits"`
	// Maximum number of write request units for the specified table.
	//
	// To specify a maximum ``OnDemandThroughput`` on your table, set the value of ``MaxWriteRequestUnits`` as greater than or equal to 1. To remove the maximum ``OnDemandThroughput`` that is currently set on your table, set the value of ``MaxWriteRequestUnits`` to -1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#max_write_request_units DynamodbTable#max_write_request_units}
	MaxWriteRequestUnits *float64 `field:"optional" json:"maxWriteRequestUnits" yaml:"maxWriteRequestUnits"`
}

