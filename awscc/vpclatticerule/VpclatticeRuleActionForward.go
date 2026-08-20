// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticerule


type VpclatticeRuleActionForward struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/vpclattice_rule#target_groups VpclatticeRule#target_groups}.
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
}

