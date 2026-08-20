// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointDashManifestsFilterConfiguration struct {
	// <p>Optionally specify the clip start time for all of your manifest egress requests.
	//
	// When you include clip start time, note that you cannot use clip start time query parameters for this manifest's endpoint URL.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#clip_start_time Mediapackagev2OriginEndpoint#clip_start_time}
	ClipStartTime *string `field:"optional" json:"clipStartTime" yaml:"clipStartTime"`
	// <p>Optionally specify one or more DRM settings for all of your manifest egress requests.
	//
	// When you include a DRM setting, note that you cannot use an identical DRM setting query parameter for this manifest's endpoint URL.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#drm_settings Mediapackagev2OriginEndpoint#drm_settings}
	DrmSettings *string `field:"optional" json:"drmSettings" yaml:"drmSettings"`
	// <p>Optionally specify the end time for all of your manifest egress requests.
	//
	// When you include end time, note that you cannot use end time query parameters for this manifest's endpoint URL.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#end Mediapackagev2OriginEndpoint#end}
	End *string `field:"optional" json:"end" yaml:"end"`
	// <p>Optionally specify one or more manifest filters for all of your manifest egress requests.
	//
	// When you include a manifest filter, note that you cannot use an identical manifest filter query parameter for this manifest's endpoint URL.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#manifest_filter Mediapackagev2OriginEndpoint#manifest_filter}
	ManifestFilter *string `field:"optional" json:"manifestFilter" yaml:"manifestFilter"`
	// <p>Optionally specify the start time for all of your manifest egress requests.
	//
	// When you include start time, note that you cannot use start time query parameters for this manifest's endpoint URL.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#start Mediapackagev2OriginEndpoint#start}
	Start *string `field:"optional" json:"start" yaml:"start"`
	// <p>Optionally specify the time delay for all of your manifest egress requests.
	//
	// Enter a value that is smaller than your endpoint's startover window. When you include time delay, note that you cannot use time delay query parameters for this manifest's endpoint URL.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#time_delay_seconds Mediapackagev2OriginEndpoint#time_delay_seconds}
	TimeDelaySeconds *float64 `field:"optional" json:"timeDelaySeconds" yaml:"timeDelaySeconds"`
}

