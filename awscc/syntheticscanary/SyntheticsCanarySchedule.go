// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanarySchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/synthetics_canary#expression SyntheticsCanary#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/synthetics_canary#duration_in_seconds SyntheticsCanary#duration_in_seconds}.
	DurationInSeconds *string `field:"optional" json:"durationInSeconds" yaml:"durationInSeconds"`
	// Provide canary auto retry configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/synthetics_canary#retry_config SyntheticsCanary#retry_config}
	RetryConfig *SyntheticsCanaryScheduleRetryConfig `field:"optional" json:"retryConfig" yaml:"retryConfig"`
}

