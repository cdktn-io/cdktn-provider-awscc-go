// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorchannel


type MediatailorChannelOutputs struct {
	// <p>The name of the manifest for the channel. The name appears in the <code>PlaybackUrl</code>.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediatailor_channel#manifest_name MediatailorChannel#manifest_name}
	ManifestName *string `field:"required" json:"manifestName" yaml:"manifestName"`
	// <p>A string used to match which <code>HttpPackageConfiguration</code> is used for each <code>VodSource</code>.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediatailor_channel#source_group MediatailorChannel#source_group}
	SourceGroup *string `field:"required" json:"sourceGroup" yaml:"sourceGroup"`
	// <p>Dash manifest configuration parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediatailor_channel#dash_playlist_settings MediatailorChannel#dash_playlist_settings}
	DashPlaylistSettings *MediatailorChannelOutputsDashPlaylistSettings `field:"optional" json:"dashPlaylistSettings" yaml:"dashPlaylistSettings"`
	// <p>HLS playlist configuration parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediatailor_channel#hls_playlist_settings MediatailorChannel#hls_playlist_settings}
	HlsPlaylistSettings *MediatailorChannelOutputsHlsPlaylistSettings `field:"optional" json:"hlsPlaylistSettings" yaml:"hlsPlaylistSettings"`
}

