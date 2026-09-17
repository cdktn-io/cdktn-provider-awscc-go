// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigdeployment


type AppconfigDeploymentDynamicExtensionParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_deployment#extension_reference AppconfigDeployment#extension_reference}.
	ExtensionReference *string `field:"optional" json:"extensionReference" yaml:"extensionReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_deployment#parameter_name AppconfigDeployment#parameter_name}.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_deployment#parameter_value AppconfigDeployment#parameter_value}.
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

