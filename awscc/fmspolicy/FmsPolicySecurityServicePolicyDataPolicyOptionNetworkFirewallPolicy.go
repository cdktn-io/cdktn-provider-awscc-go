// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy struct {
	// Firewall deployment mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fms_policy#firewall_deployment_model FmsPolicy#firewall_deployment_model}
	FirewallDeploymentModel *string `field:"optional" json:"firewallDeploymentModel" yaml:"firewallDeploymentModel"`
}

