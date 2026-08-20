// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticerule


type VpclatticeRuleMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/vpclattice_rule#http_match VpclatticeRule#http_match}.
	HttpMatch *VpclatticeRuleMatchHttpMatch `field:"required" json:"httpMatch" yaml:"httpMatch"`
}

