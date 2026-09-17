// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRulesFilter struct {
	// A container for specifying rule filters.
	//
	// The filters determine the subset of objects to which the rule applies. This element is required only if you specify more than one filter. For example:
	//   +  If you specify both a ``Prefix`` and a ``TagFilter``, wrap these filters in an ``And`` tag.
	//   +  If you specify a filter based on multiple tags, wrap the ``TagFilter`` elements in an ``And`` tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#and S3Bucket#and}
	And *S3BucketReplicationConfigurationRulesFilterAnd `field:"optional" json:"and" yaml:"and"`
	// An object key name prefix that identifies the subset of objects to which the rule applies.
	//
	// Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// A container for specifying a tag key and value.
	//
	// The rule applies only to objects that have the tag in their tag set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#tag_filter S3Bucket#tag_filter}
	TagFilter *S3BucketReplicationConfigurationRulesFilterTagFilter `field:"optional" json:"tagFilter" yaml:"tagFilter"`
}

