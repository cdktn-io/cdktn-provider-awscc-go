// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3accesspoint


type S3AccessPointTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_access_point#key S3AccessPoint#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_access_point#value S3AccessPoint#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

