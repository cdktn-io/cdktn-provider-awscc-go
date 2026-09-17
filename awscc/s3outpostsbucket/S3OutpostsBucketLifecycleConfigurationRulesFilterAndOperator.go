// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3outpostsbucket


type S3OutpostsBucketLifecycleConfigurationRulesFilterAndOperator struct {
	// Prefix identifies one or more objects to which the rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3outposts_bucket#prefix S3OutpostsBucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// All of these tags must exist in the object's tag set in order for the rule to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3outposts_bucket#tags S3OutpostsBucket#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

