// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmultiregionendpoint


type SesMultiRegionEndpointDetails struct {
	// A list of route configuration details. Must contain exactly one route configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_multi_region_endpoint#route_details SesMultiRegionEndpoint#route_details}
	RouteDetails interface{} `field:"required" json:"routeDetails" yaml:"routeDetails"`
}

