// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationlambdahook


type CloudformationLambdaHookTargetFiltersTargets struct {
	// Target actions are the type of operation hooks will be executed at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_lambda_hook#action CloudformationLambdaHook#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Invocation points are the point in provisioning workflow where hooks will be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_lambda_hook#invocation_point CloudformationLambdaHook#invocation_point}
	InvocationPoint *string `field:"optional" json:"invocationPoint" yaml:"invocationPoint"`
	// Type name of hook target. Hook targets are the destination where hooks will be invoked against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_lambda_hook#target_name CloudformationLambdaHook#target_name}
	TargetName *string `field:"optional" json:"targetName" yaml:"targetName"`
}

