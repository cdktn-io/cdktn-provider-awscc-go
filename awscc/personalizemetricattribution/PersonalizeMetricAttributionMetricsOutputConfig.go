// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizemetricattribution


type PersonalizeMetricAttributionMetricsOutputConfig struct {
	// The ARN of the IAM role for the metric attribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/personalize_metric_attribution#role_arn PersonalizeMetricAttribution#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The configuration details of an Amazon S3 output bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/personalize_metric_attribution#s3_data_destination PersonalizeMetricAttribution#s3_data_destination}
	S3DataDestination *PersonalizeMetricAttributionMetricsOutputConfigS3DataDestination `field:"optional" json:"s3DataDestination" yaml:"s3DataDestination"`
}

