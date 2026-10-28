// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination


type AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationFirehoseStream struct {
	// The name of the Amazon Kinesis Data Firehose delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_ingestion_destination#stream_name AppfabricIngestionDestination#stream_name}
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

