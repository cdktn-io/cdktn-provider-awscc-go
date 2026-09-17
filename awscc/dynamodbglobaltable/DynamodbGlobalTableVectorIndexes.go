// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableVectorIndexes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_global_table#dimensions DynamodbGlobalTable#dimensions}.
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_global_table#distance_function DynamodbGlobalTable#distance_function}.
	DistanceFunction *string `field:"optional" json:"distanceFunction" yaml:"distanceFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_global_table#index_name DynamodbGlobalTable#index_name}.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_global_table#projection DynamodbGlobalTable#projection}.
	Projection *DynamodbGlobalTableVectorIndexesProjection `field:"optional" json:"projection" yaml:"projection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_global_table#search_schema DynamodbGlobalTable#search_schema}.
	SearchSchema interface{} `field:"optional" json:"searchSchema" yaml:"searchSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_global_table#vector_attribute DynamodbGlobalTable#vector_attribute}.
	VectorAttribute *DynamodbGlobalTableVectorIndexesVectorAttribute `field:"optional" json:"vectorAttribute" yaml:"vectorAttribute"`
}

