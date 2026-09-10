// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorsourcelocation


type MediatailorSourceLocationSegmentDeliveryConfigurations struct {
	// <p>The base URL of the host or path of the segment delivery server that you're using to serve segments.
	//
	// This is typically a content delivery network (CDN). The URL can be absolute or relative. To use an absolute URL include the protocol, such as <code>https://example.com/some/path</code>. To use a relative URL specify the relative path, such as <code>/some/path*</code>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_source_location#base_url MediatailorSourceLocation#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// <p>A unique identifier used to distinguish between multiple segment delivery configurations in a source location.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_source_location#name MediatailorSourceLocation#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

