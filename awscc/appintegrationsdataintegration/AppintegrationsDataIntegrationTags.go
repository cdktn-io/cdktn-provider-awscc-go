// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appintegrationsdataintegration


type AppintegrationsDataIntegrationTags struct {
	// A key to identify the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appintegrations_data_integration#key AppintegrationsDataIntegration#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Corresponding tag value for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appintegrations_data_integration#value AppintegrationsDataIntegration#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

