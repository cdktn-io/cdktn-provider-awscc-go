// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointMssManifests struct {
	// <p>Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
	//
	// </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#filter_configuration Mediapackagev2OriginEndpoint#filter_configuration}
	FilterConfiguration *Mediapackagev2OriginEndpointMssManifestsFilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#manifest_layout Mediapackagev2OriginEndpoint#manifest_layout}.
	ManifestLayout *string `field:"optional" json:"manifestLayout" yaml:"manifestLayout"`
	// <p>The name of the MSS manifest.
	//
	// This name is appended to the origin endpoint URL to create the unique path for accessing this specific MSS manifest.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#manifest_name Mediapackagev2OriginEndpoint#manifest_name}
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// <p>The duration (in seconds) of the manifest window.
	//
	// This represents the total amount of content available in the manifest at any given time.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#manifest_window_seconds Mediapackagev2OriginEndpoint#manifest_window_seconds}
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
}

