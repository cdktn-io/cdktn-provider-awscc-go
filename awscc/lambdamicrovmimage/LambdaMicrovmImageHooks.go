// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmimage


type LambdaMicrovmImageHooks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_microvm_image#microvm_hooks LambdaMicrovmImage#microvm_hooks}.
	MicrovmHooks *LambdaMicrovmImageHooksMicrovmHooks `field:"optional" json:"microvmHooks" yaml:"microvmHooks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_microvm_image#microvm_image_hooks LambdaMicrovmImage#microvm_image_hooks}.
	MicrovmImageHooks *LambdaMicrovmImageHooksMicrovmImageHooks `field:"optional" json:"microvmImageHooks" yaml:"microvmImageHooks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_microvm_image#port LambdaMicrovmImage#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

