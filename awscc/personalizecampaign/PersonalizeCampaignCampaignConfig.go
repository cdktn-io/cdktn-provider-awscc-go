// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizecampaign


type PersonalizeCampaignCampaignConfig struct {
	// Whether metadata with recommendations is enabled for the campaign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/personalize_campaign#enable_metadata_with_recommendations PersonalizeCampaign#enable_metadata_with_recommendations}
	EnableMetadataWithRecommendations interface{} `field:"optional" json:"enableMetadataWithRecommendations" yaml:"enableMetadataWithRecommendations"`
	// Specifies the exploration configuration hyperparameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/personalize_campaign#item_exploration_config PersonalizeCampaign#item_exploration_config}
	ItemExplorationConfig *map[string]*string `field:"optional" json:"itemExplorationConfig" yaml:"itemExplorationConfig"`
	// A map of ranking influence values for POPULARITY and FRESHNESS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/personalize_campaign#ranking_influence PersonalizeCampaign#ranking_influence}
	RankingInfluence *map[string]*float64 `field:"optional" json:"rankingInfluence" yaml:"rankingInfluence"`
	// Whether the campaign automatically updates to use the latest solution version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/personalize_campaign#sync_with_latest_solution_version PersonalizeCampaign#sync_with_latest_solution_version}
	SyncWithLatestSolutionVersion interface{} `field:"optional" json:"syncWithLatestSolutionVersion" yaml:"syncWithLatestSolutionVersion"`
}

