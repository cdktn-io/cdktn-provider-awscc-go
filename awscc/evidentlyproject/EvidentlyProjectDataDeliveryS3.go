// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlyproject


type EvidentlyProjectDataDeliveryS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evidently_project#bucket_name EvidentlyProject#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evidently_project#prefix EvidentlyProject#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

