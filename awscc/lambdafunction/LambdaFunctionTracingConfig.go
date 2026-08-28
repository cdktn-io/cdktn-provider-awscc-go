// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionTracingConfig struct {
	// The tracing mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_function#mode LambdaFunction#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

