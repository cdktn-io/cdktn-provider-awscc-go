// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowSourceMediaStreamSourceConfigurations struct {
	// The format that was used to encode the data.
	//
	// For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video, 2110 streams, set the encoding name to raw. For video, JPEG XS streams, set the encoding name to jxsv.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow#encoding_name MediaconnectFlow#encoding_name}
	EncodingName *string `field:"optional" json:"encodingName" yaml:"encodingName"`
	// The media streams that you want to associate with the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow#input_configurations MediaconnectFlow#input_configurations}
	InputConfigurations interface{} `field:"optional" json:"inputConfigurations" yaml:"inputConfigurations"`
	// A name that helps you distinguish one media stream from another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow#media_stream_name MediaconnectFlow#media_stream_name}
	MediaStreamName *string `field:"optional" json:"mediaStreamName" yaml:"mediaStreamName"`
}

