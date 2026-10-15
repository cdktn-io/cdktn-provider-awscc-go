// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketObjectLockConfigurationRuleDefaultRetentionDefaultEventHold struct {
	// The number of days that you want to specify for the default event hold duration.
	//
	// Valid values are ``1`` to ``36500`` days. You must specify either ``Days`` or ``Years``, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/s3_bucket#days S3Bucket#days}
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// The number of years that you want to specify for the default event hold duration.
	//
	// Valid values are ``1`` to ``100`` years. You must specify either ``Days`` or ``Years``, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/s3_bucket#years S3Bucket#years}
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

