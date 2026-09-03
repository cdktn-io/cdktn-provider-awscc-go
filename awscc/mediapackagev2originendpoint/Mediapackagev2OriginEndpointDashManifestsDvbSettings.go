// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointDashManifestsDvbSettings struct {
	// <p>Playback device error reporting settings.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediapackagev2_origin_endpoint#error_metrics Mediapackagev2OriginEndpoint#error_metrics}
	ErrorMetrics interface{} `field:"optional" json:"errorMetrics" yaml:"errorMetrics"`
	// <p>For use with DVB-DASH profiles only.
	//
	// The settings for font downloads that you want Elemental MediaPackage to pass through to the manifest.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediapackagev2_origin_endpoint#font_download Mediapackagev2OriginEndpoint#font_download}
	FontDownload *Mediapackagev2OriginEndpointDashManifestsDvbSettingsFontDownload `field:"optional" json:"fontDownload" yaml:"fontDownload"`
}

