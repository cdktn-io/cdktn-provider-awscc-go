// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticerule


type VpclatticeRuleMatchHttpMatchHeaderMatches struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/vpclattice_rule#case_sensitive VpclatticeRule#case_sensitive}.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/vpclattice_rule#match VpclatticeRule#match}.
	Match *VpclatticeRuleMatchHttpMatchHeaderMatchesMatch `field:"optional" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/vpclattice_rule#name VpclatticeRule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

