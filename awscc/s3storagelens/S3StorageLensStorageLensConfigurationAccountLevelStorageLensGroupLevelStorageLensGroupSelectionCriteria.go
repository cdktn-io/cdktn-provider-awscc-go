// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens


type S3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevelStorageLensGroupSelectionCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_storage_lens#exclude S3StorageLens#exclude}.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_storage_lens#include S3StorageLens#include}.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

