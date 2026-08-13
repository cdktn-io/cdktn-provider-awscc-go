// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamDatabaseSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#columns KinesisfirehoseDeliveryStream#columns}.
	Columns *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationColumns `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#databases KinesisfirehoseDeliveryStream#databases}.
	Databases *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabases `field:"optional" json:"databases" yaml:"databases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#database_source_authentication_configuration KinesisfirehoseDeliveryStream#database_source_authentication_configuration}.
	DatabaseSourceAuthenticationConfiguration *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceAuthenticationConfiguration `field:"optional" json:"databaseSourceAuthenticationConfiguration" yaml:"databaseSourceAuthenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#database_source_vpc_configuration KinesisfirehoseDeliveryStream#database_source_vpc_configuration}.
	DatabaseSourceVpcConfiguration *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceVpcConfiguration `field:"optional" json:"databaseSourceVpcConfiguration" yaml:"databaseSourceVpcConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#digest KinesisfirehoseDeliveryStream#digest}.
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#endpoint KinesisfirehoseDeliveryStream#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#port KinesisfirehoseDeliveryStream#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#public_certificate KinesisfirehoseDeliveryStream#public_certificate}.
	PublicCertificate *string `field:"optional" json:"publicCertificate" yaml:"publicCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#snapshot_watermark_table KinesisfirehoseDeliveryStream#snapshot_watermark_table}.
	SnapshotWatermarkTable *string `field:"optional" json:"snapshotWatermarkTable" yaml:"snapshotWatermarkTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#ssl_mode KinesisfirehoseDeliveryStream#ssl_mode}.
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#surrogate_keys KinesisfirehoseDeliveryStream#surrogate_keys}.
	SurrogateKeys *[]*string `field:"optional" json:"surrogateKeys" yaml:"surrogateKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#tables KinesisfirehoseDeliveryStream#tables}.
	Tables *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationTables `field:"optional" json:"tables" yaml:"tables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#type KinesisfirehoseDeliveryStream#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

