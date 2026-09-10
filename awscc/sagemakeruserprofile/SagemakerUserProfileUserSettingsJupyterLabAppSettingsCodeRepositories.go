// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsJupyterLabAppSettingsCodeRepositories struct {
	// A CodeRepository (valid URL) to be used within Jupyter's Git extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_user_profile#repository_url SagemakerUserProfile#repository_url}
	RepositoryUrl *string `field:"optional" json:"repositoryUrl" yaml:"repositoryUrl"`
}

