// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstack


type AppstreamStackContentRedirectionHostToClient struct {
	// The URLs that are allowed for redirection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#allowed_urls AppstreamStack#allowed_urls}
	AllowedUrls *[]*string `field:"optional" json:"allowedUrls" yaml:"allowedUrls"`
	// The URLs that are denied for redirection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#denied_urls AppstreamStack#denied_urls}
	DeniedUrls *[]*string `field:"optional" json:"deniedUrls" yaml:"deniedUrls"`
	// Specifies whether URL redirection is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#enabled AppstreamStack#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

