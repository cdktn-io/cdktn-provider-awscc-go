// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens


type S3StorageLensStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetrics struct {
	// Specifies whether advanced cost optimization metrics are enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#is_enabled S3StorageLens#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

