// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package caseslayout


type CasesLayoutContentBasicTopPanel struct {
	// Defines the sections within a panel or tab. Contains field groups that organize related fields together.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cases_layout#sections CasesLayout#sections}
	Sections interface{} `field:"optional" json:"sections" yaml:"sections"`
}

