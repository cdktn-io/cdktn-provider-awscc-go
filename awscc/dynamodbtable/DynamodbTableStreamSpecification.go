// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableStreamSpecification struct {
	// Creates or updates a resource-based policy document that contains the permissions for DDB resources, such as a table's streams.
	//
	// Resource-based policies let you define access permissions by specifying who has access to each resource, and the actions they are allowed to perform on each resource.
	//   When you remove the ``StreamSpecification`` property from the template, DynamoDB disables the stream but retains any attached resource policy until the stream is deleted after 24 hours. When you modify the ``StreamViewType`` property, DynamoDB creates a new stream and retains the old stream's resource policy. The old stream and its resource policy are deleted after the 24-hour retention period.
	//   In a CFNshort template, you can provide the policy in JSON or YAML format because CFNshort converts YAML to JSON before submitting it to DDB. For more information about resource-based policies, see [Using resource-based policies for](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/access-control-resource-based.html) and [Resource-based policy examples](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-examples.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#resource_policy DynamodbTable#resource_policy}
	ResourcePolicy *DynamodbTableStreamSpecificationResourcePolicy `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// When an item in the table is modified, ``StreamViewType`` determines what information is written to the stream for this table.
	//
	// Valid values for ``StreamViewType`` are:
	//   +  ``KEYS_ONLY`` - Only the key attributes of the modified item are written to the stream.
	//   +  ``NEW_IMAGE`` - The entire item, as it appears after it was modified, is written to the stream.
	//   +  ``OLD_IMAGE`` - The entire item, as it appeared before it was modified, is written to the stream.
	//   +  ``NEW_AND_OLD_IMAGES`` - Both the new and the old item images of the item are written to the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#stream_view_type DynamodbTable#stream_view_type}
	StreamViewType *string `field:"optional" json:"streamViewType" yaml:"streamViewType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#tags DynamodbTable#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

