// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#block_size_bytes KinesisfirehoseDeliveryStream#block_size_bytes}.
	BlockSizeBytes *float64 `field:"optional" json:"blockSizeBytes" yaml:"blockSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#compression KinesisfirehoseDeliveryStream#compression}.
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#enable_dictionary_compression KinesisfirehoseDeliveryStream#enable_dictionary_compression}.
	EnableDictionaryCompression interface{} `field:"optional" json:"enableDictionaryCompression" yaml:"enableDictionaryCompression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#max_padding_bytes KinesisfirehoseDeliveryStream#max_padding_bytes}.
	MaxPaddingBytes *float64 `field:"optional" json:"maxPaddingBytes" yaml:"maxPaddingBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#page_size_bytes KinesisfirehoseDeliveryStream#page_size_bytes}.
	PageSizeBytes *float64 `field:"optional" json:"pageSizeBytes" yaml:"pageSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#writer_version KinesisfirehoseDeliveryStream#writer_version}.
	WriterVersion *string `field:"optional" json:"writerVersion" yaml:"writerVersion"`
}

