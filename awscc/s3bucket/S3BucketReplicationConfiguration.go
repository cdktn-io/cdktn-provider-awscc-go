// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfiguration struct {
	// The Amazon Resource Name (ARN) of the IAMlong (IAM) role that Amazon S3 assumes when replicating objects.
	//
	// For more information, see [How to Set Up Replication](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-how-setup.html) in the *Amazon S3 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#role S3Bucket#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// A container for one or more replication rules.
	//
	// A replication configuration must have at least one rule and can contain a maximum of 1,000 rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#rules S3Bucket#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

