// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmultiregionendpoint


type SesMultiRegionEndpointTags struct {
	// One part of a key-value pair that defines a tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_multi_region_endpoint#key SesMultiRegionEndpoint#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The optional part of a key-value pair that defines a tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_multi_region_endpoint#value SesMultiRegionEndpoint#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

