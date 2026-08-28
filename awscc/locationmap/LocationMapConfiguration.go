// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package locationmap


type LocationMapConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_map#style LocationMap#style}.
	Style *string `field:"required" json:"style" yaml:"style"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_map#custom_layers LocationMap#custom_layers}.
	CustomLayers *[]*string `field:"optional" json:"customLayers" yaml:"customLayers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_map#political_view LocationMap#political_view}.
	PoliticalView *string `field:"optional" json:"politicalView" yaml:"politicalView"`
}

