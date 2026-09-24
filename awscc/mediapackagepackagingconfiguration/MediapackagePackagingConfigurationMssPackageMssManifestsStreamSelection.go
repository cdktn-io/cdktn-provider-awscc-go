// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationMssPackageMssManifestsStreamSelection struct {
	// The maximum video bitrate (bps) to include in output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediapackage_packaging_configuration#max_video_bits_per_second MediapackagePackagingConfiguration#max_video_bits_per_second}
	MaxVideoBitsPerSecond *float64 `field:"optional" json:"maxVideoBitsPerSecond" yaml:"maxVideoBitsPerSecond"`
	// The minimum video bitrate (bps) to include in output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediapackage_packaging_configuration#min_video_bits_per_second MediapackagePackagingConfiguration#min_video_bits_per_second}
	MinVideoBitsPerSecond *float64 `field:"optional" json:"minVideoBitsPerSecond" yaml:"minVideoBitsPerSecond"`
	// A directive that determines the order of streams in the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediapackage_packaging_configuration#stream_order MediapackagePackagingConfiguration#stream_order}
	StreamOrder *string `field:"optional" json:"streamOrder" yaml:"streamOrder"`
}

