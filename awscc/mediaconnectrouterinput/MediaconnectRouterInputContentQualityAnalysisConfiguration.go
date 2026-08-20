// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputContentQualityAnalysisConfiguration struct {
	// Configures the content quality analysis features for the router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_input#content_level MediaconnectRouterInput#content_level}
	ContentLevel *MediaconnectRouterInputContentQualityAnalysisConfigurationContentLevel `field:"optional" json:"contentLevel" yaml:"contentLevel"`
}

