// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizesolution


type PersonalizeSolutionSolutionConfigHpoConfig struct {
	// The hyperparameters and their allowable ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/personalize_solution#algorithm_hyper_parameter_ranges PersonalizeSolution#algorithm_hyper_parameter_ranges}
	AlgorithmHyperParameterRanges *PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRanges `field:"optional" json:"algorithmHyperParameterRanges" yaml:"algorithmHyperParameterRanges"`
	// The metric to optimize during HPO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/personalize_solution#hpo_objective PersonalizeSolution#hpo_objective}
	HpoObjective *PersonalizeSolutionSolutionConfigHpoConfigHpoObjective `field:"optional" json:"hpoObjective" yaml:"hpoObjective"`
	// Describes the resource configuration for hyperparameter optimization (HPO).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/personalize_solution#hpo_resource_config PersonalizeSolution#hpo_resource_config}
	HpoResourceConfig *PersonalizeSolutionSolutionConfigHpoConfigHpoResourceConfig `field:"optional" json:"hpoResourceConfig" yaml:"hpoResourceConfig"`
}

