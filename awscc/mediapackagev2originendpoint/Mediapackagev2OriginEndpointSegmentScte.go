// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointSegmentScte struct {
	// <p>A list of additional non-Ad SCTE-35 event types to treat as advertisements.
	//
	// When configured, events matching these types produce ad markers (such as <code>SCTE35-OUT</code> and <code>SCTE35-IN</code> in HLS DATERANGE tags) in manifests.</p> <p>Valid values: <code>PROGRAM</code> | <code>CHAPTER</code> | <code>UNSCHEDULED_EVENT</code> | <code>ALTERNATE_CONTENT_OPPORTUNITY</code> | <code>NETWORK</code> </p> <p>If you don't specify any values, the default is empty (only default ad types are used).</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediapackagev2_origin_endpoint#custom_ad_types Mediapackagev2OriginEndpoint#custom_ad_types}
	CustomAdTypes *[]*string `field:"optional" json:"customAdTypes" yaml:"customAdTypes"`
	// <p>The SCTE-35 message types that you want to be treated as ad markers in the output.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediapackagev2_origin_endpoint#scte_filter Mediapackagev2OriginEndpoint#scte_filter}
	ScteFilter *[]*string `field:"optional" json:"scteFilter" yaml:"scteFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediapackagev2_origin_endpoint#scte_in_segments Mediapackagev2OriginEndpoint#scte_in_segments}.
	ScteInSegments *string `field:"optional" json:"scteInSegments" yaml:"scteInSegments"`
}

