// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset


type DatabrewDatasetInputS3InputDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/databrew_dataset#bucket DatabrewDataset#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Bucket owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/databrew_dataset#bucket_owner DatabrewDataset#bucket_owner}
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/databrew_dataset#key DatabrewDataset#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

