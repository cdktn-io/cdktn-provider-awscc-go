// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableLocalSecondaryIndexesProjection struct {
	// Represents the non-key attribute names which will be projected into the index.
	//
	// For global and local secondary indexes, the total count of ``NonKeyAttributes`` summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total. This limit only applies when you specify the ProjectionType of ``INCLUDE``. You still can specify the ProjectionType of ``ALL`` to project all attributes from the source table, even if the table has more than 100 attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_table#non_key_attributes DynamodbTable#non_key_attributes}
	NonKeyAttributes *[]*string `field:"optional" json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the index:   +  ``KEYS_ONLY`` - Only the index and primary keys are projected into the index.
	//
	// +  ``INCLUDE`` - In addition to the attributes described in ``KEYS_ONLY``, the secondary index will include other non-key attributes that you specify.
	//   +  ``ALL`` - All of the table attributes are projected into the index.
	//
	//  When using the DynamoDB console, ``ALL`` is selected by default.
	//  You can't modify the projection of an existing index. To change the projected attributes, you must delete the index and create a new one with the projection that you want.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_table#projection_type DynamodbTable#projection_type}
	ProjectionType *string `field:"optional" json:"projectionType" yaml:"projectionType"`
}

