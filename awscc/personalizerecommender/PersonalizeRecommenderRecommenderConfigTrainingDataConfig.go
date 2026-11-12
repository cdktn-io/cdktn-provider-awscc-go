// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizerecommender


type PersonalizeRecommenderRecommenderConfigTrainingDataConfig struct {
	// Specifies the columns to exclude from training.
	//
	// Each key is a dataset type, and each value is a list of columns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/personalize_recommender#excluded_dataset_columns PersonalizeRecommender#excluded_dataset_columns}
	ExcludedDatasetColumns interface{} `field:"optional" json:"excludedDatasetColumns" yaml:"excludedDatasetColumns"`
}

