// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigurationrecorder


type ConfigConfigurationRecorderRecordingMode struct {
	// The default recording frequency that AWS Config uses to record configuration changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/config_configuration_recorder#recording_frequency ConfigConfigurationRecorder#recording_frequency}
	RecordingFrequency *string `field:"optional" json:"recordingFrequency" yaml:"recordingFrequency"`
	// An array of 'RecordingModeOverride' objects for you to specify your overrides for the recording mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/config_configuration_recorder#recording_mode_overrides ConfigConfigurationRecorder#recording_mode_overrides}
	RecordingModeOverrides interface{} `field:"optional" json:"recordingModeOverrides" yaml:"recordingModeOverrides"`
}

