// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3expressdirectorybucket


type S3ExpressDirectoryBucketLifecycleConfiguration struct {
	// A lifecycle rule for individual objects in an Amazon S3 Express bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3express_directory_bucket#rules S3ExpressDirectoryBucket#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

