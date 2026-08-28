// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens


type S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsSelectionCriteria struct {
	// Delimiter to divide S3 key into hierarchy of prefixes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#delimiter S3StorageLens#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Max depth of prefixes of S3 key that Amazon S3 Storage Lens will analyze.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#max_depth S3StorageLens#max_depth}
	MaxDepth *float64 `field:"optional" json:"maxDepth" yaml:"maxDepth"`
	// The minimum storage bytes threshold for the prefixes to be included in the analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#min_storage_bytes_percentage S3StorageLens#min_storage_bytes_percentage}
	MinStorageBytesPercentage *float64 `field:"optional" json:"minStorageBytesPercentage" yaml:"minStorageBytesPercentage"`
}

