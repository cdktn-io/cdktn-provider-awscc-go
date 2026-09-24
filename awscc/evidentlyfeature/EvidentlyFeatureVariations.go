// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlyfeature


type EvidentlyFeatureVariations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/evidently_feature#boolean_value EvidentlyFeature#boolean_value}.
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/evidently_feature#double_value EvidentlyFeature#double_value}.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/evidently_feature#long_value EvidentlyFeature#long_value}.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/evidently_feature#string_value EvidentlyFeature#string_value}.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/evidently_feature#variation_name EvidentlyFeature#variation_name}.
	VariationName *string `field:"optional" json:"variationName" yaml:"variationName"`
}

