// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableIcebergMetadata struct {
	// Partition specification for an Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#iceberg_partition_spec S3TablesTable#iceberg_partition_spec}
	IcebergPartitionSpec *S3TablesTableIcebergMetadataIcebergPartitionSpec `field:"optional" json:"icebergPartitionSpec" yaml:"icebergPartitionSpec"`
	// Schema definition for flat tables with primitive types only. Mutually exclusive with IcebergSchemaV2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#iceberg_schema S3TablesTable#iceberg_schema}
	IcebergSchema *S3TablesTableIcebergMetadataIcebergSchema `field:"optional" json:"icebergSchema" yaml:"icebergSchema"`
	// Schema definition that supports Apache Iceberg nested types (struct, list, map) and primitive types. Mutually exclusive with IcebergSchema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#iceberg_schema_v2 S3TablesTable#iceberg_schema_v2}
	IcebergSchemaV2 *S3TablesTableIcebergMetadataIcebergSchemaV2 `field:"optional" json:"icebergSchemaV2" yaml:"icebergSchemaV2"`
	// Sort order specification for an Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#iceberg_sort_order S3TablesTable#iceberg_sort_order}
	IcebergSortOrder *S3TablesTableIcebergMetadataIcebergSortOrder `field:"optional" json:"icebergSortOrder" yaml:"icebergSortOrder"`
	// Iceberg table properties (e.g., format-version, write.parquet.compression-codec).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table#table_properties S3TablesTable#table_properties}
	TableProperties *map[string]*string `field:"optional" json:"tableProperties" yaml:"tableProperties"`
}

