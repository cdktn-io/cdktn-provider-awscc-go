// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowoutput


type MediaconnectFlowOutputMediaStreamOutputConfigurations struct {
	// The media streams that you want to associate with the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow_output#destination_configurations MediaconnectFlowOutput#destination_configurations}
	DestinationConfigurations interface{} `field:"optional" json:"destinationConfigurations" yaml:"destinationConfigurations"`
	// The format that will be used to encode the data.
	//
	// For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video streams on sources or outputs that use the CDI protocol, set the encoding name to raw. For video streams on sources or outputs that use the ST 2110 JPEG XS protocol, set the encoding name to jxsv.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow_output#encoding_name MediaconnectFlowOutput#encoding_name}
	EncodingName *string `field:"optional" json:"encodingName" yaml:"encodingName"`
	// A collection of parameters that determine how MediaConnect will convert the content.
	//
	// These fields only apply to outputs on flows that have a CDI source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow_output#encoding_parameters MediaconnectFlowOutput#encoding_parameters}
	EncodingParameters *MediaconnectFlowOutputMediaStreamOutputConfigurationsEncodingParameters `field:"optional" json:"encodingParameters" yaml:"encodingParameters"`
	// A name that helps you distinguish one media stream from another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow_output#media_stream_name MediaconnectFlowOutput#media_stream_name}
	MediaStreamName *string `field:"optional" json:"mediaStreamName" yaml:"mediaStreamName"`
}

