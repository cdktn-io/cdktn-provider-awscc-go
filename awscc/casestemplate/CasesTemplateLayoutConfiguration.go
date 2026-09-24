// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casestemplate


type CasesTemplateLayoutConfiguration struct {
	// The unique identifier of a layout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cases_template#default_layout CasesTemplate#default_layout}
	DefaultLayout *string `field:"optional" json:"defaultLayout" yaml:"defaultLayout"`
}

