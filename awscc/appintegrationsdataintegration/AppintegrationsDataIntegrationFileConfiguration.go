// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appintegrationsdataintegration


type AppintegrationsDataIntegrationFileConfiguration struct {
	// Restrictions for what files should be pulled from the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appintegrations_data_integration#filters AppintegrationsDataIntegration#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Identifiers for the source folders to pull all files from recursively.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appintegrations_data_integration#folders AppintegrationsDataIntegration#folders}
	Folders *[]*string `field:"optional" json:"folders" yaml:"folders"`
}

