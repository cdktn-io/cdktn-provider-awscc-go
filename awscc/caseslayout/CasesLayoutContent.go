// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package caseslayout


type CasesLayoutContent struct {
	// Defines the field layout for the agent's case interface.
	//
	// Configures which fields appear in the top panel (immediately visible) and More Info tab (expandable section) of the case view, allowing customization of the agent experience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_layout#basic CasesLayout#basic}
	Basic *CasesLayoutContentBasic `field:"optional" json:"basic" yaml:"basic"`
}

