// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointDashManifestsDvbSettingsFontDownload struct {
	// <p>The <code>fontFamily</code> name for subtitles, as described in <a href="https://tech.ebu.ch/publications/tech3380">EBU-TT-D Subtitling Distribution Format</a>. </p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediapackagev2_origin_endpoint#font_family Mediapackagev2OriginEndpoint#font_family}
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
	// <p>The <code>mimeType</code> of the resource that's at the font download URL.</p> <p>For information about font MIME types, see the <a href="https://dvb.org/wp-content/uploads/2021/06/A168r4_MPEG-DASH-Profile-for-Transport-of-ISO-BMFF-Based-DVB-Services_Draft-ts_103-285-v140_November_2021.pdf">MPEG-DASH Profile for Transport of ISO BMFF Based DVB Services over IP Based Networks</a> document. </p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediapackagev2_origin_endpoint#mime_type Mediapackagev2OriginEndpoint#mime_type}
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
	// <p>The URL for downloading fonts for subtitles.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediapackagev2_origin_endpoint#url Mediapackagev2OriginEndpoint#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

