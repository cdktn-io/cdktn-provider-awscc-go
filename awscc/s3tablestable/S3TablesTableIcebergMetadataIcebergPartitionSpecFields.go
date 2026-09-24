// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableIcebergMetadataIcebergPartitionSpecFields struct {
	// The partition field ID (auto-assigned starting from 1000 if not specified).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#field_id S3TablesTable#field_id}
	FieldId *float64 `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The name of the partition field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#name S3TablesTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The source column ID to partition on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#source_id S3TablesTable#source_id}
	SourceId *float64 `field:"optional" json:"sourceId" yaml:"sourceId"`
	// The partition transform function (identity, bucket[N], truncate[N], year, month, day, hour).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#transform S3TablesTable#transform}
	Transform *string `field:"optional" json:"transform" yaml:"transform"`
}

