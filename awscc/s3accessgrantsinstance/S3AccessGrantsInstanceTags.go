// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3accessgrantsinstance


type S3AccessGrantsInstanceTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_access_grants_instance#key S3AccessGrantsInstance#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_access_grants_instance#value S3AccessGrantsInstance#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

