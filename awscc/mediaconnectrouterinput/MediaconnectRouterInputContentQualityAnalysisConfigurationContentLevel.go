// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputContentQualityAnalysisConfigurationContentLevel struct {
	// Detects black frames in the router input's source content and reports them through a CloudWatch metric, an EventBridge event, and a router input message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_input#black_frames MediaconnectRouterInput#black_frames}
	BlackFrames *MediaconnectRouterInputContentQualityAnalysisConfigurationContentLevelBlackFrames `field:"optional" json:"blackFrames" yaml:"blackFrames"`
	// Detects frozen video frames in the router input's source content and reports them through a CloudWatch metric, an EventBridge event, and a router input message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_input#frozen_frames MediaconnectRouterInput#frozen_frames}
	FrozenFrames *MediaconnectRouterInputContentQualityAnalysisConfigurationContentLevelFrozenFrames `field:"optional" json:"frozenFrames" yaml:"frozenFrames"`
	// Detects silent audio in the router input's source content and reports it through a CloudWatch metric, an EventBridge event, and a router input message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_input#silent_audio MediaconnectRouterInput#silent_audio}
	SilentAudio *MediaconnectRouterInputContentQualityAnalysisConfigurationContentLevelSilentAudio `field:"optional" json:"silentAudio" yaml:"silentAudio"`
}

