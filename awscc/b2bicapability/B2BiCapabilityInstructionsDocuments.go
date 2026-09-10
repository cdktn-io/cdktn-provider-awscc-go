// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bicapability


type B2BiCapabilityInstructionsDocuments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/b2bi_capability#bucket_name B2BiCapability#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/b2bi_capability#key B2BiCapability#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

