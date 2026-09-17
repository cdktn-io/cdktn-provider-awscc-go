// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowoutput


type MediaconnectFlowOutputMediaStreamOutputConfigurationsEncodingParameters struct {
	// A value that is used to calculate compression for an output.
	//
	// The bitrate of the output is calculated as follows: Output bitrate = (1 / compressionFactor) * (source bitrate) This property only applies to outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol. Valid values are in the range of 3.0 to 10.0, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow_output#compression_factor MediaconnectFlowOutput#compression_factor}
	CompressionFactor *float64 `field:"optional" json:"compressionFactor" yaml:"compressionFactor"`
	// A setting on the encoder that drives compression settings.
	//
	// This property only applies to video media streams associated with outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow_output#encoder_profile MediaconnectFlowOutput#encoder_profile}
	EncoderProfile *string `field:"optional" json:"encoderProfile" yaml:"encoderProfile"`
}

