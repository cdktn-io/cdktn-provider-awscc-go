// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewallpolicy


type NetworkfirewallFirewallPolicyFirewallPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall_policy#stateless_default_actions NetworkfirewallFirewallPolicy#stateless_default_actions}.
	StatelessDefaultActions *[]*string `field:"required" json:"statelessDefaultActions" yaml:"statelessDefaultActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall_policy#stateless_fragment_default_actions NetworkfirewallFirewallPolicy#stateless_fragment_default_actions}.
	StatelessFragmentDefaultActions *[]*string `field:"required" json:"statelessFragmentDefaultActions" yaml:"statelessFragmentDefaultActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall_policy#enable_tls_session_holding NetworkfirewallFirewallPolicy#enable_tls_session_holding}.
	EnableTlsSessionHolding interface{} `field:"optional" json:"enableTlsSessionHolding" yaml:"enableTlsSessionHolding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall_policy#policy_variables NetworkfirewallFirewallPolicy#policy_variables}.
	PolicyVariables *NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariables `field:"optional" json:"policyVariables" yaml:"policyVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall_policy#stateful_default_actions NetworkfirewallFirewallPolicy#stateful_default_actions}.
	StatefulDefaultActions *[]*string `field:"optional" json:"statefulDefaultActions" yaml:"statefulDefaultActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall_policy#stateful_engine_options NetworkfirewallFirewallPolicy#stateful_engine_options}.
	StatefulEngineOptions *NetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptions `field:"optional" json:"statefulEngineOptions" yaml:"statefulEngineOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall_policy#stateful_rule_group_references NetworkfirewallFirewallPolicy#stateful_rule_group_references}.
	StatefulRuleGroupReferences interface{} `field:"optional" json:"statefulRuleGroupReferences" yaml:"statefulRuleGroupReferences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall_policy#stateless_custom_actions NetworkfirewallFirewallPolicy#stateless_custom_actions}.
	StatelessCustomActions interface{} `field:"optional" json:"statelessCustomActions" yaml:"statelessCustomActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall_policy#stateless_rule_group_references NetworkfirewallFirewallPolicy#stateless_rule_group_references}.
	StatelessRuleGroupReferences interface{} `field:"optional" json:"statelessRuleGroupReferences" yaml:"statelessRuleGroupReferences"`
	// A resource ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall_policy#tls_inspection_configuration_arn NetworkfirewallFirewallPolicy#tls_inspection_configuration_arn}
	TlsInspectionConfigurationArn *string `field:"optional" json:"tlsInspectionConfigurationArn" yaml:"tlsInspectionConfigurationArn"`
}

