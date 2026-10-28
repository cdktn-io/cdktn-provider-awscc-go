// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointDashManifestsFilterConfigurationMultiview struct {
	// <p>A tile layout for a multiview channel.
	//
	// Each layout determines how many source tiles are composited into the output and how those tiles are arranged.</p> <p>The allowed values are:</p> <ul> <li> <p> <code>LAYOUT_2EH</code> ? Two tiles of equal size, arranged horizontally.</p> </li> <li> <p> <code>LAYOUT_2PL</code> ? Two tiles, with one larger primary tile.</p> </li> <li> <p> <code>LAYOUT_3EL</code> ? Three tiles of equal size, arranged in two columns.</p> </li> <li> <p> <code>LAYOUT_3PL</code> ? Three tiles, with one larger primary tile on the left and two stacked on the right.</p> </li> <li> <p> <code>LAYOUT_4E</code> ? Four tiles of equal size, arranged in a two-by-two grid.</p> </li> <li> <p> <code>LAYOUT_4PL</code> ? Four tiles, with one larger primary tile on the left and three stacked on the right.</p> </li> </ul>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/mediapackagev2_origin_endpoint#layout Mediapackagev2OriginEndpoint#layout}
	Layout *string `field:"optional" json:"layout" yaml:"layout"`
	// <p>The source channels to composite, in tile order.
	//
	// Each channel must be one of the <code>AvailableSources</code> of the channel that this origin endpoint is on, and the number of channels must equal the number of tiles in <code>Layout</code>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/mediapackagev2_origin_endpoint#sources Mediapackagev2OriginEndpoint#sources}
	Sources *[]*string `field:"optional" json:"sources" yaml:"sources"`
}

