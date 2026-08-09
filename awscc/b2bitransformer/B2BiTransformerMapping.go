// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bitransformer


type B2BiTransformerMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/b2bi_transformer#template B2BiTransformer#template}.
	Template *string `field:"optional" json:"template" yaml:"template"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/b2bi_transformer#template_language B2BiTransformer#template_language}.
	TemplateLanguage *string `field:"optional" json:"templateLanguage" yaml:"templateLanguage"`
}

