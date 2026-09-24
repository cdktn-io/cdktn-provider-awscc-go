// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableIcebergMetadataIcebergSortOrder struct {
	// List of sort fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#fields S3TablesTable#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The sort order ID (defaults to 1 if not specified, 0 is reserved for unsorted).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#order_id S3TablesTable#order_id}
	OrderId *float64 `field:"optional" json:"orderId" yaml:"orderId"`
}

