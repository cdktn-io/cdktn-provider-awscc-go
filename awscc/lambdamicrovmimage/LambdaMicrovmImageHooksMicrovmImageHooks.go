// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmimage


type LambdaMicrovmImageHooksMicrovmImageHooks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#ready LambdaMicrovmImage#ready}.
	Ready *string `field:"optional" json:"ready" yaml:"ready"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#ready_timeout_in_seconds LambdaMicrovmImage#ready_timeout_in_seconds}.
	ReadyTimeoutInSeconds *float64 `field:"optional" json:"readyTimeoutInSeconds" yaml:"readyTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#validate LambdaMicrovmImage#validate}.
	Validate *string `field:"optional" json:"validate" yaml:"validate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#validate_timeout_in_seconds LambdaMicrovmImage#validate_timeout_in_seconds}.
	ValidateTimeoutInSeconds *float64 `field:"optional" json:"validateTimeoutInSeconds" yaml:"validateTimeoutInSeconds"`
}

