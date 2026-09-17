// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticerule


type VpclatticeRuleMatchHttpMatchHeaderMatchesMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/vpclattice_rule#contains VpclatticeRule#contains}.
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/vpclattice_rule#exact VpclatticeRule#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/vpclattice_rule#prefix VpclatticeRule#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

