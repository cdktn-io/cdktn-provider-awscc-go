// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert


type CloudwatchAlertNotificationRulesTrigger struct {
	// The alert states that trigger this rule.
	//
	// A notification is sent when the alert enters any listed state. If omitted or empty, every state triggers the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#state_values CloudwatchAlert#state_values}
	StateValues *[]*string `field:"optional" json:"stateValues" yaml:"stateValues"`
}

