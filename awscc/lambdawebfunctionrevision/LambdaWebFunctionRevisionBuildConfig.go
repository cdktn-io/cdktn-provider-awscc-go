// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdawebfunctionrevision


type LambdaWebFunctionRevisionBuildConfig struct {
	// The code configuration for the revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/lambda_web_function_revision#code_config LambdaWebFunctionRevision#code_config}
	CodeConfig *LambdaWebFunctionRevisionBuildConfigCodeConfig `field:"required" json:"codeConfig" yaml:"codeConfig"`
	// The runtime configuration for the revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/lambda_web_function_revision#runtime_config LambdaWebFunctionRevision#runtime_config}
	RuntimeConfig *LambdaWebFunctionRevisionBuildConfigRuntimeConfig `field:"required" json:"runtimeConfig" yaml:"runtimeConfig"`
}

