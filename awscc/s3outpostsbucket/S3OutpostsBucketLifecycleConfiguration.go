// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3outpostsbucket


type S3OutpostsBucketLifecycleConfiguration struct {
	// A list of lifecycle rules for individual objects in an Amazon S3Outposts bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_bucket#rules S3OutpostsBucket#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

