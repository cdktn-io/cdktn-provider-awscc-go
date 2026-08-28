// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableIcebergMetadataIcebergSortOrderFields struct {
	// Sort direction (asc or desc).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3tables_table#direction S3TablesTable#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Null value ordering (nulls-first or nulls-last).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3tables_table#null_order S3TablesTable#null_order}
	NullOrder *string `field:"optional" json:"nullOrder" yaml:"nullOrder"`
	// The source column ID to sort on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3tables_table#source_id S3TablesTable#source_id}
	SourceId *float64 `field:"optional" json:"sourceId" yaml:"sourceId"`
	// The sort transform function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3tables_table#transform S3TablesTable#transform}
	Transform *string `field:"optional" json:"transform" yaml:"transform"`
}

