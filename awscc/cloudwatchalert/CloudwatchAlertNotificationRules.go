// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert


type CloudwatchAlertNotificationRules struct {
	// Where the notification is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#target CloudwatchAlert#target}
	Target *CloudwatchAlertNotificationRulesTarget `field:"optional" json:"target" yaml:"target"`
	// The alert states that trigger the notification rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#trigger CloudwatchAlert#trigger}
	Trigger *CloudwatchAlertNotificationRulesTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

