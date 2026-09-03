// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerappimageconfig


type SagemakerAppImageConfigKernelGatewayImageConfigKernelSpecs struct {
	// The display name of the kernel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_app_image_config#display_name SagemakerAppImageConfig#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the kernel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_app_image_config#name SagemakerAppImageConfig#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

