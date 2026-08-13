// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceSourceConfigurationImageRepository struct {
	// Image Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/apprunner_service#image_configuration ApprunnerService#image_configuration}
	ImageConfiguration *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
	// Image Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/apprunner_service#image_identifier ApprunnerService#image_identifier}
	ImageIdentifier *string `field:"optional" json:"imageIdentifier" yaml:"imageIdentifier"`
	// Image Repository Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/apprunner_service#image_repository_type ApprunnerService#image_repository_type}
	ImageRepositoryType *string `field:"optional" json:"imageRepositoryType" yaml:"imageRepositoryType"`
}

