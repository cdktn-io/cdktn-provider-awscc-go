// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens


type S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetrics struct {
	// Specifies whether prefix-level storage metrics are enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3_storage_lens#is_enabled S3StorageLens#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Selection criteria for prefix-level metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3_storage_lens#selection_criteria S3StorageLens#selection_criteria}
	SelectionCriteria *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsSelectionCriteria `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
}

