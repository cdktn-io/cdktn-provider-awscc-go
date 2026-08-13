// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bitransformer


type B2BiTransformerOutputConversion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/b2bi_transformer#advanced_options B2BiTransformer#advanced_options}.
	AdvancedOptions *B2BiTransformerOutputConversionAdvancedOptions `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/b2bi_transformer#format_options B2BiTransformer#format_options}.
	FormatOptions *B2BiTransformerOutputConversionFormatOptions `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/b2bi_transformer#to_format B2BiTransformer#to_format}.
	ToFormat *string `field:"optional" json:"toFormat" yaml:"toFormat"`
}

