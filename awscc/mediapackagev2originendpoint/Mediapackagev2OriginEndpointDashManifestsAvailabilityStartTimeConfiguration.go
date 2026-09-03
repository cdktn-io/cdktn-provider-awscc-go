// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointDashManifestsAvailabilityStartTimeConfiguration struct {
	// <p>The fixed availability start time for the DASH manifest, in ISO 8601 date-time format.
	//
	// The value must have hourly granularity, meaning that the minutes, seconds, and fractional seconds must be zero. The value must be on or after <code>2024-01-01T00:00:00Z</code> and must be at least 14 days before the current time.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediapackagev2_origin_endpoint#fixed_availability_start_time Mediapackagev2OriginEndpoint#fixed_availability_start_time}
	FixedAvailabilityStartTime *string `field:"optional" json:"fixedAvailabilityStartTime" yaml:"fixedAvailabilityStartTime"`
}

