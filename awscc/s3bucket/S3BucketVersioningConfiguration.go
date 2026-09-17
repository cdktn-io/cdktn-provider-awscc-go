// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketVersioningConfiguration struct {
	// The versioning state of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#status S3Bucket#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

