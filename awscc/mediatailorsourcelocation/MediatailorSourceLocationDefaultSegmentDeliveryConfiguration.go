// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorsourcelocation


type MediatailorSourceLocationDefaultSegmentDeliveryConfiguration struct {
	// <p>The hostname of the server that will be used to serve segments.
	//
	// This string must include the protocol, such as <b>https://</b>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediatailor_source_location#base_url MediatailorSourceLocation#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
}

