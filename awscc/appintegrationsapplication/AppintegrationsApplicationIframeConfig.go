// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appintegrationsapplication


type AppintegrationsApplicationIframeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appintegrations_application#allow AppintegrationsApplication#allow}.
	Allow *[]*string `field:"optional" json:"allow" yaml:"allow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appintegrations_application#sandbox AppintegrationsApplication#sandbox}.
	Sandbox *[]*string `field:"optional" json:"sandbox" yaml:"sandbox"`
}

