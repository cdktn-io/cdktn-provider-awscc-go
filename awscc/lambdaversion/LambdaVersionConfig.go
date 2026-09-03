// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaVersionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_version#function_name LambdaVersion#function_name}
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Only publish a version if the hash value matches the value that's specified.
	//
	// Use this option to avoid publishing a version if the function code has changed since you last updated it. Updates are not supported for this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_version#code_sha_256 LambdaVersion#code_sha_256}
	CodeSha256 *string `field:"optional" json:"codeSha256" yaml:"codeSha256"`
	// A description for the version to override the description in the function configuration.
	//
	// Updates are not supported for this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_version#description LambdaVersion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The scaling configuration to apply to the function, including minimum and maximum execution environment limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_version#function_scaling_config LambdaVersion#function_scaling_config}
	FunctionScalingConfig *LambdaVersionFunctionScalingConfig `field:"optional" json:"functionScalingConfig" yaml:"functionScalingConfig"`
	// Specifies a provisioned concurrency configuration for a function's version. Updates are not supported for this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_version#provisioned_concurrency_config LambdaVersion#provisioned_concurrency_config}
	ProvisionedConcurrencyConfig *LambdaVersionProvisionedConcurrencyConfig `field:"optional" json:"provisionedConcurrencyConfig" yaml:"provisionedConcurrencyConfig"`
	// Specifies the runtime management configuration of a function. Displays runtimeVersionArn only for Manual.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_version#runtime_policy LambdaVersion#runtime_policy}
	RuntimePolicy *LambdaVersionRuntimePolicy `field:"optional" json:"runtimePolicy" yaml:"runtimePolicy"`
}

