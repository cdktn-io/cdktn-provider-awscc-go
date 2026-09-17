// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelIcebergDestinationConfiguration struct {
	// Append only mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#append_only MskChannel#append_only}
	AppendOnly interface{} `field:"optional" json:"appendOnly" yaml:"appendOnly"`
	// Catalog configuration of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#catalog MskChannel#catalog}
	Catalog *MskChannelIcebergDestinationConfigurationCatalog `field:"optional" json:"catalog" yaml:"catalog"`
	// Compression codec for Iceberg table data files. Defaults to ZSTD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#compression_type MskChannel#compression_type}
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Data freshness in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#data_freshness_in_seconds MskChannel#data_freshness_in_seconds}
	DataFreshnessInSeconds *float64 `field:"optional" json:"dataFreshnessInSeconds" yaml:"dataFreshnessInSeconds"`
	// Dead letter queue S3 configuration of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#dead_letter_queue_s3 MskChannel#dead_letter_queue_s3}
	DeadLetterQueueS3 *MskChannelIcebergDestinationConfigurationDeadLetterQueueS3 `field:"optional" json:"deadLetterQueueS3" yaml:"deadLetterQueueS3"`
	// List of destination tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#destination_table_list MskChannel#destination_table_list}
	DestinationTableList interface{} `field:"optional" json:"destinationTableList" yaml:"destinationTableList"`
	// Schema evolution configuration of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#schema_evolution MskChannel#schema_evolution}
	SchemaEvolution *MskChannelIcebergDestinationConfigurationSchemaEvolution `field:"optional" json:"schemaEvolution" yaml:"schemaEvolution"`
	// The Amazon Resource Name (ARN) of an IAM role used by MSK to access the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#service_execution_role_arn MskChannel#service_execution_role_arn}
	ServiceExecutionRoleArn *string `field:"optional" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// Table creation configuration of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#table_creation MskChannel#table_creation}
	TableCreation *MskChannelIcebergDestinationConfigurationTableCreation `field:"optional" json:"tableCreation" yaml:"tableCreation"`
}

