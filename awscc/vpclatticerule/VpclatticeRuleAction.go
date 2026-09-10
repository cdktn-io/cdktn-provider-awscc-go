// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticerule


type VpclatticeRuleAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/vpclattice_rule#fixed_response VpclatticeRule#fixed_response}.
	FixedResponse *VpclatticeRuleActionFixedResponse `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/vpclattice_rule#forward VpclatticeRule#forward}.
	Forward *VpclatticeRuleActionForward `field:"optional" json:"forward" yaml:"forward"`
}

