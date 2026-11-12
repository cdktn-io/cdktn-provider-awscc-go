// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputSrtSettingsSrtCallerSources struct {
	// The decryption settings for the SRT caller source. Present only if the source has decryption enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#decryption MedialiveInput#decryption}
	Decryption *MedialiveInputSrtSettingsSrtCallerSourcesDecryption `field:"optional" json:"decryption" yaml:"decryption"`
	// The preferred latency (in milliseconds) for implementing packet loss and recovery. Packet recovery is a key feature of SRT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#minimum_latency MedialiveInput#minimum_latency}
	MinimumLatency *float64 `field:"optional" json:"minimumLatency" yaml:"minimumLatency"`
	// The IP address at the upstream system (the listener) that MediaLive (the caller) connects to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#srt_listener_address MedialiveInput#srt_listener_address}
	SrtListenerAddress *string `field:"optional" json:"srtListenerAddress" yaml:"srtListenerAddress"`
	// The port at the upstream system (the listener) that MediaLive (the caller) connects to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#srt_listener_port MedialiveInput#srt_listener_port}
	SrtListenerPort *string `field:"optional" json:"srtListenerPort" yaml:"srtListenerPort"`
	// The stream ID, if the upstream system uses this identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#stream_id MedialiveInput#stream_id}
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

