// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsstorageconfiguration


type IvsStorageConfigurationS3 struct {
	// Location (S3 bucket name) where recorded videos will be stored.
	//
	// Note that the StorageConfiguration and S3 bucket must be in the same region as the Composition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ivs_storage_configuration#bucket_name IvsStorageConfiguration#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

