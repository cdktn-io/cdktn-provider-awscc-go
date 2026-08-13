// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputContentQualityAnalysisConfigurationContentLevelFrozenFrames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_input#state MediaconnectRouterInput#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// The number of consecutive seconds of a frozen frame that MediaConnect must detect before it reports an issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_input#threshold_seconds MediaconnectRouterInput#threshold_seconds}
	ThresholdSeconds *float64 `field:"optional" json:"thresholdSeconds" yaml:"thresholdSeconds"`
}

