// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceSourceConfigurationCodeRepository struct {
	// Code Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/apprunner_service#code_configuration ApprunnerService#code_configuration}
	CodeConfiguration *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
	// Repository Url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/apprunner_service#repository_url ApprunnerService#repository_url}
	RepositoryUrl *string `field:"optional" json:"repositoryUrl" yaml:"repositoryUrl"`
	// Source Code Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/apprunner_service#source_code_version ApprunnerService#source_code_version}
	SourceCodeVersion *ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion `field:"optional" json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
	// Source Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/apprunner_service#source_directory ApprunnerService#source_directory}
	SourceDirectory *string `field:"optional" json:"sourceDirectory" yaml:"sourceDirectory"`
}

