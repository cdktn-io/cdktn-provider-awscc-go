// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorsourcelocation


type MediatailorSourceLocationHttpConfiguration struct {
	// <p>The base URL for the source location host server. This string must include the protocol, such as <b>https://</b>.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_source_location#base_url MediatailorSourceLocation#base_url}
	BaseUrl *string `field:"required" json:"baseUrl" yaml:"baseUrl"`
}

