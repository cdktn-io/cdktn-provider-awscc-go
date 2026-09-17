// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtrailtrail


type CloudtrailTrailInsightSelectors struct {
	// The categories of events for which to log insights. By default, insights are logged for management events only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#event_categories CloudtrailTrail#event_categories}
	EventCategories *[]*string `field:"optional" json:"eventCategories" yaml:"eventCategories"`
	// The type of insight to log on a trail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#insight_type CloudtrailTrail#insight_type}
	InsightType *string `field:"optional" json:"insightType" yaml:"insightType"`
}

