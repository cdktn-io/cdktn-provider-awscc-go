// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowSourceMonitoringConfigVideoMonitoringSettings struct {
	// Configures settings for the BlackFrames metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#black_frames MediaconnectFlow#black_frames}
	BlackFrames *MediaconnectFlowSourceMonitoringConfigVideoMonitoringSettingsBlackFrames `field:"optional" json:"blackFrames" yaml:"blackFrames"`
	// Configures settings for the FrozenFrames metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#frozen_frames MediaconnectFlow#frozen_frames}
	FrozenFrames *MediaconnectFlowSourceMonitoringConfigVideoMonitoringSettingsFrozenFrames `field:"optional" json:"frozenFrames" yaml:"frozenFrames"`
}

