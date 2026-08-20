// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationCmafPackageHlsManifests struct {
	// This setting controls how ad markers are included in the packaged OriginEndpoint.
	//
	// "NONE" will omit all SCTE-35 ad markers from the output. "PASSTHROUGH" causes the manifest to contain a copy of the SCTE-35 ad markers (comments) taken directly from the input HTTP Live Streaming (HLS) manifest. "SCTE35_ENHANCED" generates ad markers and blackout tags based on SCTE-35 messages in the input source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackage_packaging_configuration#ad_markers MediapackagePackagingConfiguration#ad_markers}
	AdMarkers *string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// When enabled, an I-Frame only stream will be included in the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackage_packaging_configuration#include_iframe_only_stream MediapackagePackagingConfiguration#include_iframe_only_stream}
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// An optional string to include in the name of the manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackage_packaging_configuration#manifest_name MediapackagePackagingConfiguration#manifest_name}
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME tag inserted into manifests.
	//
	// Additionally, when an interval is specified ID3Timed Metadata messages will be generated every 5 seconds using the ingest time of the content. If the interval is not specified, or set to 0, then no EXT-X-PROGRAM-DATE-TIME tags will be inserted into manifests and no ID3Timed Metadata messages will be generated. Note that irrespective of this parameter, if any ID3 Timed Metadata is found in HTTP Live Streaming (HLS) input, it will be passed through to HLS output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackage_packaging_configuration#program_date_time_interval_seconds MediapackagePackagingConfiguration#program_date_time_interval_seconds}
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// When enabled, the EXT-X-KEY tag will be repeated in output manifests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackage_packaging_configuration#repeat_ext_x_key MediapackagePackagingConfiguration#repeat_ext_x_key}
	RepeatExtXKey interface{} `field:"optional" json:"repeatExtXKey" yaml:"repeatExtXKey"`
	// A StreamSelection configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackage_packaging_configuration#stream_selection MediapackagePackagingConfiguration#stream_selection}
	StreamSelection *MediapackagePackagingConfigurationCmafPackageHlsManifestsStreamSelection `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

