// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableIcebergMetadataIcebergPartitionSpec struct {
	// List of partition fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#fields S3TablesTable#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The partition spec ID (defaults to 0 if not specified).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#spec_id S3TablesTable#spec_id}
	SpecId *float64 `field:"optional" json:"specId" yaml:"specId"`
}

