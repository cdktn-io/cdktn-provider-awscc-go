// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstack


type AppstreamStackStreamingExperienceSettings struct {
	// The preferred protocol that you want to use while streaming your application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appstream_stack#preferred_protocol AppstreamStack#preferred_protocol}
	PreferredProtocol *string `field:"optional" json:"preferredProtocol" yaml:"preferredProtocol"`
}

