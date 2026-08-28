// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketCorsConfiguration struct {
	// A set of origins and methods (cross-origin access that you want to allow).
	//
	// You can add up to 100 rules to the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#cors_rules S3Bucket#cors_rules}
	CorsRules interface{} `field:"optional" json:"corsRules" yaml:"corsRules"`
}

