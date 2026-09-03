// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsCustomImages struct {
	// The Name of the AppImageConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#app_image_config_name SagemakerDomain#app_image_config_name}
	AppImageConfigName *string `field:"optional" json:"appImageConfigName" yaml:"appImageConfigName"`
	// The name of the CustomImage. Must be unique to your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#image_name SagemakerDomain#image_name}
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// The version number of the CustomImage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#image_version_number SagemakerDomain#image_version_number}
	ImageVersionNumber *float64 `field:"optional" json:"imageVersionNumber" yaml:"imageVersionNumber"`
}

