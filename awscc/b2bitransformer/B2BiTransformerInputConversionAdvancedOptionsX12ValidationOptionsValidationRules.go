// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bitransformer


type B2BiTransformerInputConversionAdvancedOptionsX12ValidationOptionsValidationRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/b2bi_transformer#code_list_validation_rule B2BiTransformer#code_list_validation_rule}.
	CodeListValidationRule *B2BiTransformerInputConversionAdvancedOptionsX12ValidationOptionsValidationRulesCodeListValidationRule `field:"optional" json:"codeListValidationRule" yaml:"codeListValidationRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/b2bi_transformer#element_length_validation_rule B2BiTransformer#element_length_validation_rule}.
	ElementLengthValidationRule *B2BiTransformerInputConversionAdvancedOptionsX12ValidationOptionsValidationRulesElementLengthValidationRule `field:"optional" json:"elementLengthValidationRule" yaml:"elementLengthValidationRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/b2bi_transformer#element_requirement_validation_rule B2BiTransformer#element_requirement_validation_rule}.
	ElementRequirementValidationRule *B2BiTransformerInputConversionAdvancedOptionsX12ValidationOptionsValidationRulesElementRequirementValidationRule `field:"optional" json:"elementRequirementValidationRule" yaml:"elementRequirementValidationRule"`
}

