// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointDashManifestsDvbSettingsErrorMetrics struct {
	// <p>The number of playback devices per 1000 that will send error reports to the reporting URL.
	//
	// This represents the probability that a playback device will be a reporting player for this session.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#probability Mediapackagev2OriginEndpoint#probability}
	Probability *float64 `field:"optional" json:"probability" yaml:"probability"`
	// <p>The URL where playback devices send error reports.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint#reporting_url Mediapackagev2OriginEndpoint#reporting_url}
	ReportingUrl *string `field:"optional" json:"reportingUrl" yaml:"reportingUrl"`
}

