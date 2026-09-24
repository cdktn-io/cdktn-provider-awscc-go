// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmimage


type LambdaMicrovmImageHooksMicrovmHooks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#resume LambdaMicrovmImage#resume}.
	Resume *string `field:"optional" json:"resume" yaml:"resume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#resume_timeout_in_seconds LambdaMicrovmImage#resume_timeout_in_seconds}.
	ResumeTimeoutInSeconds *float64 `field:"optional" json:"resumeTimeoutInSeconds" yaml:"resumeTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#run LambdaMicrovmImage#run}.
	Run *string `field:"optional" json:"run" yaml:"run"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#run_timeout_in_seconds LambdaMicrovmImage#run_timeout_in_seconds}.
	RunTimeoutInSeconds *float64 `field:"optional" json:"runTimeoutInSeconds" yaml:"runTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#suspend LambdaMicrovmImage#suspend}.
	Suspend *string `field:"optional" json:"suspend" yaml:"suspend"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#suspend_timeout_in_seconds LambdaMicrovmImage#suspend_timeout_in_seconds}.
	SuspendTimeoutInSeconds *float64 `field:"optional" json:"suspendTimeoutInSeconds" yaml:"suspendTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#terminate LambdaMicrovmImage#terminate}.
	Terminate *string `field:"optional" json:"terminate" yaml:"terminate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#terminate_timeout_in_seconds LambdaMicrovmImage#terminate_timeout_in_seconds}.
	TerminateTimeoutInSeconds *float64 `field:"optional" json:"terminateTimeoutInSeconds" yaml:"terminateTimeoutInSeconds"`
}

