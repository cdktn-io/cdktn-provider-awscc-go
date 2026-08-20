// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2componentversion


type Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_component_version#container_params Greengrassv2ComponentVersion#container_params}.
	ContainerParams *Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParams `field:"optional" json:"containerParams" yaml:"containerParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_component_version#isolation_mode Greengrassv2ComponentVersion#isolation_mode}.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
}

