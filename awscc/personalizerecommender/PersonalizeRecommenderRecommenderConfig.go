// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizerecommender


type PersonalizeRecommenderRecommenderConfig struct {
	// Whether metadata with recommendations is enabled for the recommender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/personalize_recommender#enable_metadata_with_recommendations PersonalizeRecommender#enable_metadata_with_recommendations}
	EnableMetadataWithRecommendations interface{} `field:"optional" json:"enableMetadataWithRecommendations" yaml:"enableMetadataWithRecommendations"`
	// Specifies the exploration configuration hyperparameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/personalize_recommender#item_exploration_config PersonalizeRecommender#item_exploration_config}
	ItemExplorationConfig *map[string]*string `field:"optional" json:"itemExplorationConfig" yaml:"itemExplorationConfig"`
	// Specifies the requested minimum provisioned recommendation requests per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/personalize_recommender#min_recommendation_requests_per_second PersonalizeRecommender#min_recommendation_requests_per_second}
	MinRecommendationRequestsPerSecond *float64 `field:"optional" json:"minRecommendationRequestsPerSecond" yaml:"minRecommendationRequestsPerSecond"`
	// The training data configuration to use when creating a domain recommender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/personalize_recommender#training_data_config PersonalizeRecommender#training_data_config}
	TrainingDataConfig *PersonalizeRecommenderRecommenderConfigTrainingDataConfig `field:"optional" json:"trainingDataConfig" yaml:"trainingDataConfig"`
}

