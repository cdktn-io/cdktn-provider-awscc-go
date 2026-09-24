// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appintegrationseventintegration


type AppintegrationsEventIntegrationTags struct {
	// A key to identify the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appintegrations_event_integration#key AppintegrationsEventIntegration#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Corresponding tag value for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appintegrations_event_integration#value AppintegrationsEventIntegration#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

