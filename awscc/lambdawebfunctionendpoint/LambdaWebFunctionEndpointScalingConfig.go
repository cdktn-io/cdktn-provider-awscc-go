// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdawebfunctionendpoint


type LambdaWebFunctionEndpointScalingConfig struct {
	// The maximum number of concurrent execution environments for the endpoint.
	//
	// This optional limit further constrains the endpoint's scaling. When omitted, the endpoint's scaling is limited only by your account's vCPU quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/lambda_web_function_endpoint#max_environments LambdaWebFunctionEndpoint#max_environments}
	MaxEnvironments *float64 `field:"optional" json:"maxEnvironments" yaml:"maxEnvironments"`
}

