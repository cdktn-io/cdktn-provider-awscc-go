// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointForceEndpointErrorConfiguration struct {
	// <p>The failover conditions for the endpoint.
	//
	// The options are:</p> <ul> <li> <p> <code>STALE_MANIFEST</code> - The manifest stalled and there are no new segments or parts.</p> </li> <li> <p> <code>INCOMPLETE_MANIFEST</code> - There is a gap in the manifest.</p> </li> <li> <p> <code>MISSING_DRM_KEY</code> - Key rotation is enabled but we're unable to fetch the key for the current key period.</p> </li> <li> <p> <code>SLATE_INPUT</code> - The segments which contain slate content are considered to be missing content.</p> </li> </ul>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediapackagev2_origin_endpoint#endpoint_error_conditions Mediapackagev2OriginEndpoint#endpoint_error_conditions}
	EndpointErrorConditions *[]*string `field:"optional" json:"endpointErrorConditions" yaml:"endpointErrorConditions"`
}

