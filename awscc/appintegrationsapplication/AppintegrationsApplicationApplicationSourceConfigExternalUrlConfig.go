// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appintegrationsapplication


type AppintegrationsApplicationApplicationSourceConfigExternalUrlConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appintegrations_application#access_url AppintegrationsApplication#access_url}.
	AccessUrl *string `field:"required" json:"accessUrl" yaml:"accessUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appintegrations_application#approved_origins AppintegrationsApplication#approved_origins}.
	ApprovedOrigins *[]*string `field:"optional" json:"approvedOrigins" yaml:"approvedOrigins"`
}

