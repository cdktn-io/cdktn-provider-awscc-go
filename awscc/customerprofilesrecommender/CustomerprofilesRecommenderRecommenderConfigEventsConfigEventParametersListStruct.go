// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesrecommender


type CustomerprofilesRecommenderRecommenderConfigEventsConfigEventParametersListStruct struct {
	// The type of event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_recommender#event_type CustomerprofilesRecommender#event_type}
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// The threshold of the event type.
	//
	// Only events with a value greater or equal to this threshold will be considered for solution creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_recommender#event_value_threshold CustomerprofilesRecommender#event_value_threshold}
	EventValueThreshold *float64 `field:"optional" json:"eventValueThreshold" yaml:"eventValueThreshold"`
}

