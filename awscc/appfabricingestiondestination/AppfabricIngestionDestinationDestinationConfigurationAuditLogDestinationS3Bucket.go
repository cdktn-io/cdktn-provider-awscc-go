// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination


type AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3Bucket struct {
	// The name of the Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_ingestion_destination#bucket_name AppfabricIngestionDestination#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The object key to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_ingestion_destination#prefix AppfabricIngestionDestination#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

