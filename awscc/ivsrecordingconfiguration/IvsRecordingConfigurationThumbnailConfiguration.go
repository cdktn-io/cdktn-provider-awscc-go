// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsrecordingconfiguration


type IvsRecordingConfigurationThumbnailConfiguration struct {
	// Thumbnail Recording Mode, which determines whether thumbnails are recorded at an interval or are disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_recording_configuration#recording_mode IvsRecordingConfiguration#recording_mode}
	RecordingMode *string `field:"optional" json:"recordingMode" yaml:"recordingMode"`
	// Resolution indicates the desired resolution of recorded thumbnails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_recording_configuration#resolution IvsRecordingConfiguration#resolution}
	Resolution *string `field:"optional" json:"resolution" yaml:"resolution"`
	// Storage indicates the format in which thumbnails are recorded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_recording_configuration#storage IvsRecordingConfiguration#storage}
	Storage *[]*string `field:"optional" json:"storage" yaml:"storage"`
	// Target Interval Seconds defines the interval at which thumbnails are recorded. This field is required if RecordingMode is INTERVAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_recording_configuration#target_interval_seconds IvsRecordingConfiguration#target_interval_seconds}
	TargetIntervalSeconds *float64 `field:"optional" json:"targetIntervalSeconds" yaml:"targetIntervalSeconds"`
}

