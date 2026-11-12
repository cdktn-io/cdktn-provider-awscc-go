// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert


type CloudwatchAlertNotificationRulesTarget struct {
	// The ARN of the target.
	//
	// For sns, the ARN of the Amazon SNS topic. For slack and pagerduty, the ARN of a CloudWatch Omni integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#arn CloudwatchAlert#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Settings specific to the target, as key-value pairs (up to 20).
	//
	// The keys depend on the integration type. Not used for sns targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#metadata CloudwatchAlert#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// The target type. Valid values: sns, slack, pagerduty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#type CloudwatchAlert#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

