// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Mediapackagev2OriginEndpointConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#channel_group_name Mediapackagev2OriginEndpoint#channel_group_name}.
	ChannelGroupName *string `field:"required" json:"channelGroupName" yaml:"channelGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#channel_name Mediapackagev2OriginEndpoint#channel_name}.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#container_type Mediapackagev2OriginEndpoint#container_type}.
	ContainerType *string `field:"required" json:"containerType" yaml:"containerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#origin_endpoint_name Mediapackagev2OriginEndpoint#origin_endpoint_name}.
	OriginEndpointName *string `field:"required" json:"originEndpointName" yaml:"originEndpointName"`
	// <p>A DASH manifest configuration.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#dash_manifests Mediapackagev2OriginEndpoint#dash_manifests}
	DashManifests interface{} `field:"optional" json:"dashManifests" yaml:"dashManifests"`
	// <p>Enter any descriptive text that helps you to identify the origin endpoint.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#description Mediapackagev2OriginEndpoint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// <p>The failover settings for the endpoint.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#force_endpoint_error_configuration Mediapackagev2OriginEndpoint#force_endpoint_error_configuration}
	ForceEndpointErrorConfiguration *Mediapackagev2OriginEndpointForceEndpointErrorConfiguration `field:"optional" json:"forceEndpointErrorConfiguration" yaml:"forceEndpointErrorConfiguration"`
	// <p>An HTTP live streaming (HLS) manifest configuration.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#hls_manifests Mediapackagev2OriginEndpoint#hls_manifests}
	HlsManifests interface{} `field:"optional" json:"hlsManifests" yaml:"hlsManifests"`
	// <p>A low-latency HLS manifest configuration.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#low_latency_hls_manifests Mediapackagev2OriginEndpoint#low_latency_hls_manifests}
	LowLatencyHlsManifests interface{} `field:"optional" json:"lowLatencyHlsManifests" yaml:"lowLatencyHlsManifests"`
	// <p>The Microsoft Smooth Streaming (MSS) manifest configurations associated with this origin endpoint.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#mss_manifests Mediapackagev2OriginEndpoint#mss_manifests}
	MssManifests interface{} `field:"optional" json:"mssManifests" yaml:"mssManifests"`
	// <p>The segment configuration, including the segment name, duration, and other configuration values.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#segment Mediapackagev2OriginEndpoint#segment}
	Segment *Mediapackagev2OriginEndpointSegment `field:"optional" json:"segment" yaml:"segment"`
	// <p>The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing.
	//
	// Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#startover_window_seconds Mediapackagev2OriginEndpoint#startover_window_seconds}
	StartoverWindowSeconds *float64 `field:"optional" json:"startoverWindowSeconds" yaml:"startoverWindowSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#tags Mediapackagev2OriginEndpoint#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackagev2_origin_endpoint#uri_separator Mediapackagev2OriginEndpoint#uri_separator}.
	UriSeparator *string `field:"optional" json:"uriSeparator" yaml:"uriSeparator"`
}

