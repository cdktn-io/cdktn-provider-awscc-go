// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens


type S3StorageLensStorageLensConfigurationAccountLevelBucketLevel struct {
	// Enables activity metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_storage_lens#activity_metrics S3StorageLens#activity_metrics}
	ActivityMetrics *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelActivityMetrics `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// Enables advanced cost optimization metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_storage_lens#advanced_cost_optimization_metrics S3StorageLens#advanced_cost_optimization_metrics}
	AdvancedCostOptimizationMetrics *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetrics `field:"optional" json:"advancedCostOptimizationMetrics" yaml:"advancedCostOptimizationMetrics"`
	// Enables advanced data protection metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_storage_lens#advanced_data_protection_metrics S3StorageLens#advanced_data_protection_metrics}
	AdvancedDataProtectionMetrics *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetrics `field:"optional" json:"advancedDataProtectionMetrics" yaml:"advancedDataProtectionMetrics"`
	// Advanced Performance Metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_storage_lens#advanced_performance_metrics S3StorageLens#advanced_performance_metrics}
	AdvancedPerformanceMetrics *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetrics `field:"optional" json:"advancedPerformanceMetrics" yaml:"advancedPerformanceMetrics"`
	// Enables detailed status codes metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_storage_lens#detailed_status_codes_metrics S3StorageLens#detailed_status_codes_metrics}
	DetailedStatusCodesMetrics *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodesMetrics `field:"optional" json:"detailedStatusCodesMetrics" yaml:"detailedStatusCodesMetrics"`
	// Prefix-level metrics configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_storage_lens#prefix_level S3StorageLens#prefix_level}
	PrefixLevel *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevel `field:"optional" json:"prefixLevel" yaml:"prefixLevel"`
}

