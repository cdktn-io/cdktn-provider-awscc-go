// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableStreamSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_global_table#stream_view_type DynamodbGlobalTable#stream_view_type}.
	StreamViewType *string `field:"optional" json:"streamViewType" yaml:"streamViewType"`
}

