// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableVectorIndexes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_table#dimensions DynamodbTable#dimensions}.
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_table#distance_function DynamodbTable#distance_function}.
	DistanceFunction *string `field:"optional" json:"distanceFunction" yaml:"distanceFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_table#index_name DynamodbTable#index_name}.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Represents attributes that are copied (projected) from the table into an index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_table#projection DynamodbTable#projection}
	Projection *DynamodbTableVectorIndexesProjection `field:"optional" json:"projection" yaml:"projection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_table#search_schema DynamodbTable#search_schema}.
	SearchSchema interface{} `field:"optional" json:"searchSchema" yaml:"searchSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_table#vector_attribute DynamodbTable#vector_attribute}.
	VectorAttribute *DynamodbTableVectorIndexesVectorAttribute `field:"optional" json:"vectorAttribute" yaml:"vectorAttribute"`
}

