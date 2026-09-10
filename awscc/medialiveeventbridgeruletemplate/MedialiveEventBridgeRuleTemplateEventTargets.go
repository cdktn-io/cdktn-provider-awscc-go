// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveeventbridgeruletemplate


type MedialiveEventBridgeRuleTemplateEventTargets struct {
	// Target ARNs must be either an SNS topic or CloudWatch log group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/medialive_event_bridge_rule_template#arn MedialiveEventBridgeRuleTemplate#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

