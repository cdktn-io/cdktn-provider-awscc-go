// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration struct {
	// Code Configuration Values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apprunner_service#code_configuration_values ApprunnerService#code_configuration_values}
	CodeConfigurationValues *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues `field:"optional" json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
	// Configuration Source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apprunner_service#configuration_source ApprunnerService#configuration_source}
	ConfigurationSource *string `field:"optional" json:"configurationSource" yaml:"configurationSource"`
}

