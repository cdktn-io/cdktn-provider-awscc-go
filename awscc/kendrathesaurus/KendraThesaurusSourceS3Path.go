// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendrathesaurus


type KendraThesaurusSourceS3Path struct {
	// The name of the S3 bucket that contains the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_thesaurus#bucket KendraThesaurus#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_thesaurus#key KendraThesaurus#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

