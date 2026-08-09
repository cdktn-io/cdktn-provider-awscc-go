// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bitransformer


type B2BiTransformerInputConversionAdvancedOptionsX12ValidationOptionsValidationRulesCodeListValidationRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/b2bi_transformer#codes_to_add B2BiTransformer#codes_to_add}.
	CodesToAdd *[]*string `field:"optional" json:"codesToAdd" yaml:"codesToAdd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/b2bi_transformer#codes_to_remove B2BiTransformer#codes_to_remove}.
	CodesToRemove *[]*string `field:"optional" json:"codesToRemove" yaml:"codesToRemove"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/b2bi_transformer#element_id B2BiTransformer#element_id}.
	ElementId *string `field:"optional" json:"elementId" yaml:"elementId"`
}

