// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination


type AppfabricIngestionDestinationDestinationConfigurationAuditLogDestination struct {
	// Contains information about an Amazon Kinesis Data Firehose delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/appfabric_ingestion_destination#firehose_stream AppfabricIngestionDestination#firehose_stream}
	FirehoseStream *AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationFirehoseStream `field:"optional" json:"firehoseStream" yaml:"firehoseStream"`
	// Contains information about an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/appfabric_ingestion_destination#s3_bucket AppfabricIngestionDestination#s3_bucket}
	S3Bucket *AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3Bucket `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

