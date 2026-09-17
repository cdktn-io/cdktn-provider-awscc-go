// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableGlobalSecondaryIndexesProvisionedThroughput struct {
	// The maximum number of strongly consistent reads consumed per second before DynamoDB returns a ``ThrottlingException``.
	//
	// For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html) in the *Amazon DynamoDB Developer Guide*.
	//  If read/write capacity mode is ``PAY_PER_REQUEST`` the value is set to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_table#read_capacity_units DynamodbTable#read_capacity_units}
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// The maximum number of writes consumed per second before DynamoDB returns a ``ThrottlingException``.
	//
	// For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html) in the *Amazon DynamoDB Developer Guide*.
	//  If read/write capacity mode is ``PAY_PER_REQUEST`` the value is set to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_table#write_capacity_units DynamodbTable#write_capacity_units}
	WriteCapacityUnits *float64 `field:"optional" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

