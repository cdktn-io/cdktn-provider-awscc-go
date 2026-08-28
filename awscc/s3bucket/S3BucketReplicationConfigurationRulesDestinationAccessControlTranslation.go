// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRulesDestinationAccessControlTranslation struct {
	// Specifies the replica ownership. For default and valid values, see [PUT bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) in the *Amazon S3 API Reference*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#owner S3Bucket#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
}

