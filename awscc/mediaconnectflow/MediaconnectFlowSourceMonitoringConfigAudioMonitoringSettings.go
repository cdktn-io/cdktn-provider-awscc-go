// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettings struct {
	// Configures settings for the SilentAudio metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#silent_audio MediaconnectFlow#silent_audio}
	SilentAudio *MediaconnectFlowSourceMonitoringConfigAudioMonitoringSettingsSilentAudio `field:"optional" json:"silentAudio" yaml:"silentAudio"`
}

