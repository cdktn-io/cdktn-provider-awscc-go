// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bitransformer


type B2BiTransformerInputConversionAdvancedOptionsX12 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/b2bi_transformer#split_options B2BiTransformer#split_options}.
	SplitOptions *B2BiTransformerInputConversionAdvancedOptionsX12SplitOptions `field:"optional" json:"splitOptions" yaml:"splitOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/b2bi_transformer#validation_options B2BiTransformer#validation_options}.
	ValidationOptions *B2BiTransformerInputConversionAdvancedOptionsX12ValidationOptions `field:"optional" json:"validationOptions" yaml:"validationOptions"`
}

