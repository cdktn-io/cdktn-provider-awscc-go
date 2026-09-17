// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewallpolicy


type NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferences struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkfirewall_firewall_policy#deep_threat_inspection NetworkfirewallFirewallPolicy#deep_threat_inspection}.
	DeepThreatInspection interface{} `field:"optional" json:"deepThreatInspection" yaml:"deepThreatInspection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkfirewall_firewall_policy#override NetworkfirewallFirewallPolicy#override}.
	Override *NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferencesOverride `field:"optional" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkfirewall_firewall_policy#priority NetworkfirewallFirewallPolicy#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A resource ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkfirewall_firewall_policy#resource_arn NetworkfirewallFirewallPolicy#resource_arn}
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

