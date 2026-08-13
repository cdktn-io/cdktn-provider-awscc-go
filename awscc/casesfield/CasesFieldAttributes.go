// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casesfield


type CasesFieldAttributes struct {
	// Field attributes for Text field type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cases_field#text CasesField#text}
	Text *CasesFieldAttributesText `field:"optional" json:"text" yaml:"text"`
}

