// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketAccelerateConfiguration struct {
	// Specifies the transfer acceleration status of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#acceleration_status S3Bucket#acceleration_status}
	AccelerationStatus *string `field:"optional" json:"accelerationStatus" yaml:"accelerationStatus"`
}

