// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistener


type VpclatticeListenerDefaultAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/vpclattice_listener#fixed_response VpclatticeListener#fixed_response}.
	FixedResponse *VpclatticeListenerDefaultActionFixedResponse `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/vpclattice_listener#forward VpclatticeListener#forward}.
	Forward *VpclatticeListenerDefaultActionForward `field:"optional" json:"forward" yaml:"forward"`
}

