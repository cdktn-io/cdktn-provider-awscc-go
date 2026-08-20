// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketObjectLockConfigurationRuleDefaultRetention struct {
	// The number of days that you want to specify for the default retention period.
	//
	// If Object Lock is turned on, you must specify ``Mode`` and specify either ``Days`` or ``Years``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#days S3Bucket#days}
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// The default Object Lock retention mode you want to apply to new objects placed in the specified bucket.
	//
	// If Object Lock is turned on, you must specify ``Mode`` and specify either ``Days`` or ``Years``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#mode S3Bucket#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The number of years that you want to specify for the default retention period.
	//
	// If Object Lock is turned on, you must specify ``Mode`` and specify either ``Days`` or ``Years``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#years S3Bucket#years}
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

