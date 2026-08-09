// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmultiregionendpoint


type SesMultiRegionEndpointDetailsRouteDetails struct {
	// The name of an AWS-Region to be a secondary region for the multi-region endpoint (global-endpoint).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_multi_region_endpoint#region SesMultiRegionEndpoint#region}
	Region *string `field:"required" json:"region" yaml:"region"`
}

