// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filesaccesspoint


type S3FilesAccessPointTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3files_access_point#key S3FilesAccessPoint#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3files_access_point#value S3FilesAccessPoint#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

