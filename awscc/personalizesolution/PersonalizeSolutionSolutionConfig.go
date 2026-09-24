// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizesolution


type PersonalizeSolutionSolutionConfig struct {
	// Lists the hyperparameter names and ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/personalize_solution#algorithm_hyper_parameters PersonalizeSolution#algorithm_hyper_parameters}
	AlgorithmHyperParameters *map[string]*string `field:"optional" json:"algorithmHyperParameters" yaml:"algorithmHyperParameters"`
	// The AutoMLConfig object containing a list of recipes to search when AutoML is performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/personalize_solution#auto_ml_config PersonalizeSolution#auto_ml_config}
	AutoMlConfig *PersonalizeSolutionSolutionConfigAutoMlConfig `field:"optional" json:"autoMlConfig" yaml:"autoMlConfig"`
	// Only events with a value greater than or equal to this threshold are used for training a model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/personalize_solution#event_value_threshold PersonalizeSolution#event_value_threshold}
	EventValueThreshold *string `field:"optional" json:"eventValueThreshold" yaml:"eventValueThreshold"`
	// Lists the feature transformation parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/personalize_solution#feature_transformation_parameters PersonalizeSolution#feature_transformation_parameters}
	FeatureTransformationParameters *map[string]*string `field:"optional" json:"featureTransformationParameters" yaml:"featureTransformationParameters"`
	// Describes the properties for hyperparameter optimization (HPO).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/personalize_solution#hpo_config PersonalizeSolution#hpo_config}
	HpoConfig *PersonalizeSolutionSolutionConfigHpoConfig `field:"optional" json:"hpoConfig" yaml:"hpoConfig"`
}

