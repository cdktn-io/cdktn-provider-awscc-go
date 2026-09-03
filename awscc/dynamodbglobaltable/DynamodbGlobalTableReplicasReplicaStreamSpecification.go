// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableReplicasReplicaStreamSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dynamodb_global_table#resource_policy DynamodbGlobalTable#resource_policy}.
	ResourcePolicy *DynamodbGlobalTableReplicasReplicaStreamSpecificationResourcePolicy `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dynamodb_global_table#tags DynamodbGlobalTable#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

