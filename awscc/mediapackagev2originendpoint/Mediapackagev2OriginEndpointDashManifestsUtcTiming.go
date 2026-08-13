// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointDashManifestsUtcTiming struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediapackagev2_origin_endpoint#timing_mode Mediapackagev2OriginEndpoint#timing_mode}.
	TimingMode *string `field:"optional" json:"timingMode" yaml:"timingMode"`
	// <p>The the method that the player uses to synchronize to coordinated universal time (UTC) wall clock time.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediapackagev2_origin_endpoint#timing_source Mediapackagev2OriginEndpoint#timing_source}
	TimingSource *string `field:"optional" json:"timingSource" yaml:"timingSource"`
}

