// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointDashManifestsBaseUrls struct {
	// <p>For use with DVB-DASH profiles only.
	//
	// The priority of this location for servings segments. The lower the number, the higher the priority.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediapackagev2_origin_endpoint#dvb_priority Mediapackagev2OriginEndpoint#dvb_priority}
	DvbPriority *float64 `field:"optional" json:"dvbPriority" yaml:"dvbPriority"`
	// <p>For use with DVB-DASH profiles only. The weighting for source locations that have the same priority. </p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediapackagev2_origin_endpoint#dvb_weight Mediapackagev2OriginEndpoint#dvb_weight}
	DvbWeight *float64 `field:"optional" json:"dvbWeight" yaml:"dvbWeight"`
	// <p>The name of the source location.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediapackagev2_origin_endpoint#service_location Mediapackagev2OriginEndpoint#service_location}
	ServiceLocation *string `field:"optional" json:"serviceLocation" yaml:"serviceLocation"`
	// <p>A source location for segments.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediapackagev2_origin_endpoint#url Mediapackagev2OriginEndpoint#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

