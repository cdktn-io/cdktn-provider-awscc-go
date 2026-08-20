// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens


type S3StorageLensStorageLensConfiguration struct {
	// Account-level metrics configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#account_level S3StorageLens#account_level}
	AccountLevel *S3StorageLensStorageLensConfigurationAccountLevel `field:"required" json:"accountLevel" yaml:"accountLevel"`
	// The ID that identifies the Amazon S3 Storage Lens configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#id S3StorageLens#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Specifies whether the Amazon S3 Storage Lens configuration is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#is_enabled S3StorageLens#is_enabled}
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// The AWS Organizations ARN to use in the Amazon S3 Storage Lens configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#aws_org S3StorageLens#aws_org}
	AwsOrg *S3StorageLensStorageLensConfigurationAwsOrg `field:"optional" json:"awsOrg" yaml:"awsOrg"`
	// Specifies how Amazon S3 Storage Lens metrics should be exported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#data_export S3StorageLens#data_export}
	DataExport *S3StorageLensStorageLensConfigurationDataExport `field:"optional" json:"dataExport" yaml:"dataExport"`
	// S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#exclude S3StorageLens#exclude}
	Exclude *S3StorageLensStorageLensConfigurationExclude `field:"optional" json:"exclude" yaml:"exclude"`
	// Expanded Prefixes Data Export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#expanded_prefixes_data_export S3StorageLens#expanded_prefixes_data_export}
	ExpandedPrefixesDataExport *S3StorageLensStorageLensConfigurationExpandedPrefixesDataExport `field:"optional" json:"expandedPrefixesDataExport" yaml:"expandedPrefixesDataExport"`
	// S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#include S3StorageLens#include}
	Include *S3StorageLensStorageLensConfigurationInclude `field:"optional" json:"include" yaml:"include"`
	// The delimiter to divide S3 key into hierarchy of prefixes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_storage_lens#prefix_delimiter S3StorageLens#prefix_delimiter}
	PrefixDelimiter *string `field:"optional" json:"prefixDelimiter" yaml:"prefixDelimiter"`
}

