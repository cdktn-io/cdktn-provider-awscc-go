// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationDashPackage struct {
	// A list of DASH manifest configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackage_packaging_configuration#dash_manifests MediapackagePackagingConfiguration#dash_manifests}
	DashManifests interface{} `field:"optional" json:"dashManifests" yaml:"dashManifests"`
	// A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackage_packaging_configuration#encryption MediapackagePackagingConfiguration#encryption}
	Encryption *MediapackagePackagingConfigurationDashPackageEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// When includeEncoderConfigurationInSegments is set to true, MediaPackage places your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and Video Parameter Set (VPS) metadata in every video segment instead of in the init fragment.
	//
	// This lets you use different SPS/PPS/VPS settings for your assets during content playback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackage_packaging_configuration#include_encoder_configuration_in_segments MediapackagePackagingConfiguration#include_encoder_configuration_in_segments}
	IncludeEncoderConfigurationInSegments interface{} `field:"optional" json:"includeEncoderConfigurationInSegments" yaml:"includeEncoderConfigurationInSegments"`
	// When enabled, an I-Frame only stream will be included in the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackage_packaging_configuration#include_iframe_only_stream MediapackagePackagingConfiguration#include_iframe_only_stream}
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// A list of triggers that controls when the outgoing Dynamic Adaptive Streaming over HTTP (DASH) Media Presentation Description (MPD) will be partitioned into multiple periods.
	//
	// If empty, the content will not be partitioned into more than one period. If the list contains "ADS", new periods will be created where the Asset contains SCTE-35 ad markers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackage_packaging_configuration#period_triggers MediapackagePackagingConfiguration#period_triggers}
	PeriodTriggers *[]*string `field:"optional" json:"periodTriggers" yaml:"periodTriggers"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments will be rounded to the nearest multiple of the source fragment duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackage_packaging_configuration#segment_duration_seconds MediapackagePackagingConfiguration#segment_duration_seconds}
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// Determines the type of SegmentTemplate included in the Media Presentation Description (MPD).
	//
	// When set to NUMBER_WITH_TIMELINE, a full timeline is presented in each SegmentTemplate, with $Number$ media URLs. When set to TIME_WITH_TIMELINE, a full timeline is presented in each SegmentTemplate, with $Time$ media URLs. When set to NUMBER_WITH_DURATION, only a duration is included in each SegmentTemplate, with $Number$ media URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackage_packaging_configuration#segment_template_format MediapackagePackagingConfiguration#segment_template_format}
	SegmentTemplateFormat *string `field:"optional" json:"segmentTemplateFormat" yaml:"segmentTemplateFormat"`
}

