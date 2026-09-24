// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codestarnotificationsnotificationrule


type CodestarnotificationsNotificationRuleTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codestarnotifications_notification_rule#target_address CodestarnotificationsNotificationRule#target_address}.
	TargetAddress *string `field:"required" json:"targetAddress" yaml:"targetAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codestarnotifications_notification_rule#target_type CodestarnotificationsNotificationRule#target_type}.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

