// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointHlsManifestsStartTag struct {
	// <p>Specify the value for PRECISE within your EXT-X-START tag.
	//
	// Leave blank, or choose false, to use the default value NO. Choose yes to use the value YES.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackagev2_origin_endpoint#precise Mediapackagev2OriginEndpoint#precise}
	Precise interface{} `field:"optional" json:"precise" yaml:"precise"`
	// <p>Specify the value for TIME-OFFSET within your EXT-X-START tag.
	//
	// Enter a signed floating point value which, if positive, must be less than the configured manifest duration minus three times the configured segment target duration. If negative, the absolute value must be larger than three times the configured segment target duration, and the absolute value must be smaller than the configured manifest duration.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackagev2_origin_endpoint#time_offset Mediapackagev2OriginEndpoint#time_offset}
	TimeOffset *float64 `field:"optional" json:"timeOffset" yaml:"timeOffset"`
}

