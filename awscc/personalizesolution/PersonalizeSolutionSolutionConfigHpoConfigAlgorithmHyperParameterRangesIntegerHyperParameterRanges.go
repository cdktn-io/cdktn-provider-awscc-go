// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizesolution


type PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesIntegerHyperParameterRanges struct {
	// The maximum allowable value for the hyperparameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/personalize_solution#max_value PersonalizeSolution#max_value}
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// The minimum allowable value for the hyperparameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/personalize_solution#min_value PersonalizeSolution#min_value}
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// The name of the hyperparameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/personalize_solution#name PersonalizeSolution#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

