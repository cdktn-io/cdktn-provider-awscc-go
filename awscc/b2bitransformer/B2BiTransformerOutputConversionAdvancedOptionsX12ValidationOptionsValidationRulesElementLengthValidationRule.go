// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bitransformer


type B2BiTransformerOutputConversionAdvancedOptionsX12ValidationOptionsValidationRulesElementLengthValidationRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/b2bi_transformer#element_id B2BiTransformer#element_id}.
	ElementId *string `field:"optional" json:"elementId" yaml:"elementId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/b2bi_transformer#max_length B2BiTransformer#max_length}.
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/b2bi_transformer#min_length B2BiTransformer#min_length}.
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
}

