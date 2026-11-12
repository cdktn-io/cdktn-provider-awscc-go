// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputSrtSettings struct {
	// Placeholder documentation for __listOfSrtCallerSource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#srt_caller_sources MedialiveInput#srt_caller_sources}
	SrtCallerSources interface{} `field:"optional" json:"srtCallerSources" yaml:"srtCallerSources"`
	// Settings for an SRT listener input. Configured only when the input uses the SRT_LISTENER input type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#srt_listener_settings MedialiveInput#srt_listener_settings}
	SrtListenerSettings *MedialiveInputSrtSettingsSrtListenerSettings `field:"optional" json:"srtListenerSettings" yaml:"srtListenerSettings"`
}

