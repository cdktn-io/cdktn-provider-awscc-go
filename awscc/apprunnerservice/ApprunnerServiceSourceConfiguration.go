// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceSourceConfiguration struct {
	// Authentication Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/apprunner_service#authentication_configuration ApprunnerService#authentication_configuration}
	AuthenticationConfiguration *ApprunnerServiceSourceConfigurationAuthenticationConfiguration `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Auto Deployment enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/apprunner_service#auto_deployments_enabled ApprunnerService#auto_deployments_enabled}
	AutoDeploymentsEnabled interface{} `field:"optional" json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// Source Code Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/apprunner_service#code_repository ApprunnerService#code_repository}
	CodeRepository *ApprunnerServiceSourceConfigurationCodeRepository `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// Image Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/apprunner_service#image_repository ApprunnerService#image_repository}
	ImageRepository *ApprunnerServiceSourceConfigurationImageRepository `field:"optional" json:"imageRepository" yaml:"imageRepository"`
}

