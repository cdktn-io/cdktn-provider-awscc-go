// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3outpostsbucket


type S3OutpostsBucketLifecycleConfigurationRulesFilter struct {
	// The container for the AND condition for the lifecycle rule.
	//
	// A combination of Prefix and 1 or more Tags OR a minimum of 2 or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_bucket#and_operator S3OutpostsBucket#and_operator}
	AndOperator *S3OutpostsBucketLifecycleConfigurationRulesFilterAndOperator `field:"optional" json:"andOperator" yaml:"andOperator"`
	// Object key prefix that identifies one or more objects to which this rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_bucket#prefix S3OutpostsBucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Specifies a tag used to identify a subset of objects for an Amazon S3Outposts bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_bucket#tag S3OutpostsBucket#tag}
	Tag *S3OutpostsBucketLifecycleConfigurationRulesFilterTag `field:"optional" json:"tag" yaml:"tag"`
}

