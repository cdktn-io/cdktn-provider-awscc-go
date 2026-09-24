// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventinvokeconfig


type LambdaEventInvokeConfigDestinationConfig struct {
	// The destination configuration for failed invocations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_event_invoke_config#on_failure LambdaEventInvokeConfig#on_failure}
	OnFailure *LambdaEventInvokeConfigDestinationConfigOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination configuration for successful invocations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_event_invoke_config#on_success LambdaEventInvokeConfig#on_success}
	OnSuccess *LambdaEventInvokeConfigDestinationConfigOnSuccess `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

