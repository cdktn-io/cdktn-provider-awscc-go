// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesrecommender


type CustomerprofilesRecommenderRecommenderConfig struct {
	// Configuration for events used in the recommender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/customerprofiles_recommender#events_config CustomerprofilesRecommender#events_config}
	EventsConfig *CustomerprofilesRecommenderRecommenderConfigEventsConfig `field:"optional" json:"eventsConfig" yaml:"eventsConfig"`
}

