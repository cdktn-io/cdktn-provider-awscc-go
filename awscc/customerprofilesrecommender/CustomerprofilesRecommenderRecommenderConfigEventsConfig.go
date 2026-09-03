// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesrecommender


type CustomerprofilesRecommenderRecommenderConfigEventsConfig struct {
	// List of event parameters with their value thresholds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_recommender#event_parameters_list CustomerprofilesRecommender#event_parameters_list}
	EventParametersList interface{} `field:"optional" json:"eventParametersList" yaml:"eventParametersList"`
}

