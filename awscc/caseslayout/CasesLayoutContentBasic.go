// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package caseslayout


type CasesLayoutContentBasic struct {
	// Sections within a panel or tab of the page layout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_layout#more_info CasesLayout#more_info}
	MoreInfo *CasesLayoutContentBasicMoreInfo `field:"optional" json:"moreInfo" yaml:"moreInfo"`
	// Sections within a panel or tab of the page layout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_layout#top_panel CasesLayout#top_panel}
	TopPanel *CasesLayoutContentBasicTopPanel `field:"optional" json:"topPanel" yaml:"topPanel"`
}

