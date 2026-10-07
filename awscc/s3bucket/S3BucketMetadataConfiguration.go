// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketMetadataConfiguration struct {
	// The annotation table configuration for a metadata configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#annotation_table_configuration S3Bucket#annotation_table_configuration}
	AnnotationTableConfiguration *S3BucketMetadataConfigurationAnnotationTableConfiguration `field:"optional" json:"annotationTableConfiguration" yaml:"annotationTableConfiguration"`
	// The inventory table configuration for a metadata configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#inventory_table_configuration S3Bucket#inventory_table_configuration}
	InventoryTableConfiguration *S3BucketMetadataConfigurationInventoryTableConfiguration `field:"optional" json:"inventoryTableConfiguration" yaml:"inventoryTableConfiguration"`
	// The journal table configuration for a metadata configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#journal_table_configuration S3Bucket#journal_table_configuration}
	JournalTableConfiguration *S3BucketMetadataConfigurationJournalTableConfiguration `field:"optional" json:"journalTableConfiguration" yaml:"journalTableConfiguration"`
}

