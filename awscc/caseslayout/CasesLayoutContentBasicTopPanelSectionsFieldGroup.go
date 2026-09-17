// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package caseslayout


type CasesLayoutContentBasicTopPanelSectionsFieldGroup struct {
	// An ordered list of fields to display in this group.
	//
	// The order determines the sequence in which fields appear in the agent interface. Each field is referenced by its unique field ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cases_layout#fields CasesLayout#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// A descriptive name for the field group. Helps organize related fields together in the layout interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cases_layout#name CasesLayout#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

