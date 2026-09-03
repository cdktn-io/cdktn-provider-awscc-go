// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdacapacityprovider


type LambdaCapacityProviderPropagateTags struct {
	// A list of tags to explicitly propagate to managed resources. Maximum of 40 tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_capacity_provider#explicit_tags LambdaCapacityProvider#explicit_tags}
	ExplicitTags interface{} `field:"optional" json:"explicitTags" yaml:"explicitTags"`
	// The mode for tag propagation. Use ``Explicit`` to propagate specific tags, or ``None`` to disable propagation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_capacity_provider#mode LambdaCapacityProvider#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

