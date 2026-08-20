// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointSegment struct {
	// <p>The parameters for encrypting content.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#encryption Mediapackagev2OriginEndpoint#encryption}
	Encryption *Mediapackagev2OriginEndpointSegmentEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// <p>When selected, the stream set includes an additional I-frame only stream, along with the other tracks.
	//
	// If false, this extra stream is not included. MediaPackage generates an I-frame only stream from the first rendition in the manifest. The service inserts EXT-I-FRAMES-ONLY tags in the output manifest, and then generates and includes an I-frames only playlist in the stream. This playlist permits player functionality like fast forward and rewind.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#include_iframe_only_streams Mediapackagev2OriginEndpoint#include_iframe_only_streams}
	IncludeIframeOnlyStreams interface{} `field:"optional" json:"includeIframeOnlyStreams" yaml:"includeIframeOnlyStreams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#output_timestamp_mode Mediapackagev2OriginEndpoint#output_timestamp_mode}.
	OutputTimestampMode *string `field:"optional" json:"outputTimestampMode" yaml:"outputTimestampMode"`
	// <p>The SCTE configuration.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#scte Mediapackagev2OriginEndpoint#scte}
	Scte *Mediapackagev2OriginEndpointSegmentScte `field:"optional" json:"scte" yaml:"scte"`
	// <p>The duration (in seconds) of each segment.
	//
	// Enter a value equal to, or a multiple of, the input segment duration. If the value that you enter is different from the input segment duration, MediaPackage rounds segments to the nearest multiple of the input segment duration.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#segment_duration_seconds Mediapackagev2OriginEndpoint#segment_duration_seconds}
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// <p>The name that describes the segment.
	//
	// The name is the base name of the segment used in all content manifests inside of the endpoint. You can't use spaces in the name.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#segment_name Mediapackagev2OriginEndpoint#segment_name}
	SegmentName *string `field:"optional" json:"segmentName" yaml:"segmentName"`
	// <p>By default, MediaPackage excludes all digital video broadcasting (DVB) subtitles from the output.
	//
	// When selected, MediaPackage passes through DVB subtitles into the output.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#ts_include_dvb_subtitles Mediapackagev2OriginEndpoint#ts_include_dvb_subtitles}
	TsIncludeDvbSubtitles interface{} `field:"optional" json:"tsIncludeDvbSubtitles" yaml:"tsIncludeDvbSubtitles"`
	// <p>When selected, MediaPackage bundles all audio tracks in a rendition group.
	//
	// All other tracks in the stream can be used with any audio rendition from the group.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#ts_use_audio_rendition_group Mediapackagev2OriginEndpoint#ts_use_audio_rendition_group}
	TsUseAudioRenditionGroup interface{} `field:"optional" json:"tsUseAudioRenditionGroup" yaml:"tsUseAudioRenditionGroup"`
}

