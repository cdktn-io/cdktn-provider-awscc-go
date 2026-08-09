// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsStudioWebPortalSettingsHiddenSageMakerImageVersionAliases struct {
	// The SageMaker image name that you are hiding from the Studio user interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_user_profile#sage_maker_image_name SagemakerUserProfile#sage_maker_image_name}
	SageMakerImageName *string `field:"optional" json:"sageMakerImageName" yaml:"sageMakerImageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_user_profile#version_aliases SagemakerUserProfile#version_aliases}.
	VersionAliases *[]*string `field:"optional" json:"versionAliases" yaml:"versionAliases"`
}

