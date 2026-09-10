// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimeappinstancebot


type ChimeAppInstanceBotConfigurationLexInvokedBy struct {
	// Sets standard messages as the bot trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/chime_app_instance_bot#standard_messages ChimeAppInstanceBot#standard_messages}
	StandardMessages *string `field:"optional" json:"standardMessages" yaml:"standardMessages"`
	// Sets targeted messages as the bot trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/chime_app_instance_bot#targeted_messages ChimeAppInstanceBot#targeted_messages}
	TargetedMessages *string `field:"optional" json:"targetedMessages" yaml:"targetedMessages"`
}

