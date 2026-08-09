// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRulesDestinationMetricsEventThreshold struct {
	// Contains an integer specifying time in minutes.    Valid value: 15.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_bucket#minutes S3Bucket#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
}

