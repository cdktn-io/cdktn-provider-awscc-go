// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizesolution


type PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRangesCategoricalHyperParameterRanges struct {
	// The name of the hyperparameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/personalize_solution#name PersonalizeSolution#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the categories for the hyperparameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/personalize_solution#values PersonalizeSolution#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

