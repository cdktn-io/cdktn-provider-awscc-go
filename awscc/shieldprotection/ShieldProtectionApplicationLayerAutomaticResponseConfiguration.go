// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package shieldprotection


type ShieldProtectionApplicationLayerAutomaticResponseConfiguration struct {
	// Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks.
	//
	// You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/shield_protection#action ShieldProtection#action}
	Action *ShieldProtectionApplicationLayerAutomaticResponseConfigurationAction `field:"optional" json:"action" yaml:"action"`
	// Indicates whether automatic application layer DDoS mitigation is enabled for the protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/shield_protection#status ShieldProtection#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

