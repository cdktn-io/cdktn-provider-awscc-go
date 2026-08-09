// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtrailtrail


type CloudtrailTrailAggregationConfigurations struct {
	// The category of events to be aggregated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudtrail_trail#event_category CloudtrailTrail#event_category}
	EventCategory *string `field:"optional" json:"eventCategory" yaml:"eventCategory"`
	// Contains all templates in an aggregation configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudtrail_trail#templates CloudtrailTrail#templates}
	Templates *[]*string `field:"optional" json:"templates" yaml:"templates"`
}

