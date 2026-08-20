// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3accessgrantslocation


type S3AccessGrantsLocationTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_access_grants_location#key S3AccessGrantsLocation#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_access_grants_location#value S3AccessGrantsLocation#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

