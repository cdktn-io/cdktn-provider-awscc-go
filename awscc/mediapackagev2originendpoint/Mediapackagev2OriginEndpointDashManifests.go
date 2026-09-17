// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointDashManifests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#audio_timeline_pattern Mediapackagev2OriginEndpoint#audio_timeline_pattern}.
	AudioTimelinePattern *string `field:"optional" json:"audioTimelinePattern" yaml:"audioTimelinePattern"`
	// <p>The configuration for the DASH <code>availabilityStartTime</code> attribute of the Media Presentation Description (MPD).
	//
	// Use this configuration to set a custom availability start time for your DASH manifest.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#availability_start_time_configuration Mediapackagev2OriginEndpoint#availability_start_time_configuration}
	AvailabilityStartTimeConfiguration *Mediapackagev2OriginEndpointDashManifestsAvailabilityStartTimeConfiguration `field:"optional" json:"availabilityStartTimeConfiguration" yaml:"availabilityStartTimeConfiguration"`
	// <p>The base URL to use for retrieving segments.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#base_urls Mediapackagev2OriginEndpoint#base_urls}
	BaseUrls interface{} `field:"optional" json:"baseUrls" yaml:"baseUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#compactness Mediapackagev2OriginEndpoint#compactness}.
	Compactness *string `field:"optional" json:"compactness" yaml:"compactness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#drm_signaling Mediapackagev2OriginEndpoint#drm_signaling}.
	DrmSignaling *string `field:"optional" json:"drmSignaling" yaml:"drmSignaling"`
	// <p>For endpoints that use the DVB-DASH profile only.
	//
	// The font download and error reporting information that you want MediaPackage to pass through to the manifest.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#dvb_settings Mediapackagev2OriginEndpoint#dvb_settings}
	DvbSettings *Mediapackagev2OriginEndpointDashManifestsDvbSettings `field:"optional" json:"dvbSettings" yaml:"dvbSettings"`
	// <p>Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
	//
	// </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#filter_configuration Mediapackagev2OriginEndpoint#filter_configuration}
	FilterConfiguration *Mediapackagev2OriginEndpointDashManifestsFilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// <p>A short string that's appended to the endpoint URL.
	//
	// The manifest name creates a unique path to this endpoint. If you don't enter a value, MediaPackage uses the default manifest name, index. </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#manifest_name Mediapackagev2OriginEndpoint#manifest_name}
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// <p>The total duration (in seconds) of the manifest's content.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#manifest_window_seconds Mediapackagev2OriginEndpoint#manifest_window_seconds}
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// <p>Minimum amount of content (in seconds) that a player must keep available in the buffer.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#min_buffer_time_seconds Mediapackagev2OriginEndpoint#min_buffer_time_seconds}
	MinBufferTimeSeconds *float64 `field:"optional" json:"minBufferTimeSeconds" yaml:"minBufferTimeSeconds"`
	// <p>Minimum amount of time (in seconds) that the player should wait before requesting updates to the manifest.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#min_update_period_seconds Mediapackagev2OriginEndpoint#min_update_period_seconds}
	MinUpdatePeriodSeconds *float64 `field:"optional" json:"minUpdatePeriodSeconds" yaml:"minUpdatePeriodSeconds"`
	// <p>A list of triggers that controls when AWS Elemental MediaPackage separates the MPEG-DASH manifest into multiple periods.
	//
	// Leave this value empty to indicate that the manifest is contained all in one period. For more information about periods in the DASH manifest, see <a href="https://docs.aws.amazon.com/mediapackage/latest/userguide/multi-period.html">Multi-period DASH in AWS Elemental MediaPackage</a>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#period_triggers Mediapackagev2OriginEndpoint#period_triggers}
	PeriodTriggers *[]*string `field:"optional" json:"periodTriggers" yaml:"periodTriggers"`
	// <p>The profile that the output is compliant with.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#profiles Mediapackagev2OriginEndpoint#profiles}
	Profiles *[]*string `field:"optional" json:"profiles" yaml:"profiles"`
	// <p>Details about the content that you want MediaPackage to pass through in the manifest to the playback device.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#program_information Mediapackagev2OriginEndpoint#program_information}
	ProgramInformation *Mediapackagev2OriginEndpointDashManifestsProgramInformation `field:"optional" json:"programInformation" yaml:"programInformation"`
	// <p>The SCTE configuration.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#scte_dash Mediapackagev2OriginEndpoint#scte_dash}
	ScteDash *Mediapackagev2OriginEndpointDashManifestsScteDash `field:"optional" json:"scteDash" yaml:"scteDash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#segment_template_format Mediapackagev2OriginEndpoint#segment_template_format}.
	SegmentTemplateFormat *string `field:"optional" json:"segmentTemplateFormat" yaml:"segmentTemplateFormat"`
	// <p>The configuration for DASH subtitles.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#subtitle_configuration Mediapackagev2OriginEndpoint#subtitle_configuration}
	SubtitleConfiguration *Mediapackagev2OriginEndpointDashManifestsSubtitleConfiguration `field:"optional" json:"subtitleConfiguration" yaml:"subtitleConfiguration"`
	// <p>The amount of time (in seconds) that the player should be from the end of the manifest.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#suggested_presentation_delay_seconds Mediapackagev2OriginEndpoint#suggested_presentation_delay_seconds}
	SuggestedPresentationDelaySeconds *float64 `field:"optional" json:"suggestedPresentationDelaySeconds" yaml:"suggestedPresentationDelaySeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#uri_path_type Mediapackagev2OriginEndpoint#uri_path_type}.
	UriPathType *string `field:"optional" json:"uriPathType" yaml:"uriPathType"`
	// <p>Determines the type of UTC timing included in the DASH Media Presentation Description (MPD).</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#utc_timing Mediapackagev2OriginEndpoint#utc_timing}
	UtcTiming *Mediapackagev2OriginEndpointDashManifestsUtcTiming `field:"optional" json:"utcTiming" yaml:"utcTiming"`
}

