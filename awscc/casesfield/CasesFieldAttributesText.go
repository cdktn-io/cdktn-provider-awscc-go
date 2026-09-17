// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casesfield


type CasesFieldAttributesText struct {
	// Attribute that defines rendering component and validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cases_field#is_multiline CasesField#is_multiline}
	IsMultiline interface{} `field:"optional" json:"isMultiline" yaml:"isMultiline"`
}

