// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsencoderconfiguration


type IvsEncoderConfigurationVideo struct {
	// Bitrate for generated output, in bps. Default: 2500000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ivs_encoder_configuration#bitrate IvsEncoderConfiguration#bitrate}
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Video frame rate, in fps. Default: 30.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ivs_encoder_configuration#framerate IvsEncoderConfiguration#framerate}
	Framerate *float64 `field:"optional" json:"framerate" yaml:"framerate"`
	// Video-resolution height.
	//
	// This must be an even number. Note that the maximum value is determined by width times height, such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 720.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ivs_encoder_configuration#height IvsEncoderConfiguration#height}
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Video-resolution width.
	//
	// This must be an even number. Note that the maximum value is determined by width times height, such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 1280.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ivs_encoder_configuration#width IvsEncoderConfiguration#width}
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

