// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelPrimaryContainerImageConfig struct {
	// Set this to one of the following values: Platform - The model image is hosted in Amazon ECR.
	//
	// Vpc - The model image is hosted in a private Docker registry in your VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_model#repository_access_mode SagemakerModel#repository_access_mode}
	RepositoryAccessMode *string `field:"optional" json:"repositoryAccessMode" yaml:"repositoryAccessMode"`
	// Specifies an authentication configuration for the private docker registry where your model image is hosted.
	//
	// Specify a value for this property only if you specified `Vpc` as the value for the `RepositoryAccessMode` field of the `ImageConfig` object that you passed to a call to `CreateModel` and the private Docker registry where the model image is hosted requires authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_model#repository_auth_config SagemakerModel#repository_auth_config}
	RepositoryAuthConfig *SagemakerModelPrimaryContainerImageConfigRepositoryAuthConfig `field:"optional" json:"repositoryAuthConfig" yaml:"repositoryAuthConfig"`
}

