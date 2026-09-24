// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2channel


type Mediapackagev2ChannelMultiviewConfiguration struct {
	// <p>The tile layouts that players can request from this multiview channel's origin endpoints.
	//
	// Only the layouts that you list here are available. Each layout must appear at most once.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediapackagev2_channel#available_layouts Mediapackagev2Channel#available_layouts}
	AvailableLayouts *[]*string `field:"optional" json:"availableLayouts" yaml:"availableLayouts"`
	// <p>The channels that players can use as tiles in this multiview channel's output.
	//
	// Each source channel must be in the same channel group as the multiview channel, and must have an <code>InputType</code> of <code>CMAF</code>. Only the channels that you list here are available as tiles.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediapackagev2_channel#available_sources Mediapackagev2Channel#available_sources}
	AvailableSources *[]*string `field:"optional" json:"availableSources" yaml:"availableSources"`
}

