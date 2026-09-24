// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowSourceMonitoringConfig struct {
	// Contains the settings for audio stream metrics monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#audio_monitoring_settings MediaconnectFlow#audio_monitoring_settings}
	AudioMonitoringSettings interface{} `field:"optional" json:"audioMonitoringSettings" yaml:"audioMonitoringSettings"`
	// Indicates whether content quality analysis is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#content_quality_analysis_state MediaconnectFlow#content_quality_analysis_state}
	ContentQualityAnalysisState *string `field:"optional" json:"contentQualityAnalysisState" yaml:"contentQualityAnalysisState"`
	// The state of thumbnail monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#thumbnail_state MediaconnectFlow#thumbnail_state}
	ThumbnailState *string `field:"optional" json:"thumbnailState" yaml:"thumbnailState"`
	// Contains the settings for video stream metrics monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#video_monitoring_settings MediaconnectFlow#video_monitoring_settings}
	VideoMonitoringSettings interface{} `field:"optional" json:"videoMonitoringSettings" yaml:"videoMonitoringSettings"`
}

