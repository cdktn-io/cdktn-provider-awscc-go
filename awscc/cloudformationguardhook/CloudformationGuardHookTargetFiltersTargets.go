// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationguardhook


type CloudformationGuardHookTargetFiltersTargets struct {
	// Target actions are the type of operation hooks will be executed at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_guard_hook#action CloudformationGuardHook#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Invocation points are the point in provisioning workflow where hooks will be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_guard_hook#invocation_point CloudformationGuardHook#invocation_point}
	InvocationPoint *string `field:"optional" json:"invocationPoint" yaml:"invocationPoint"`
	// Type name of hook target. Hook targets are the destination where hooks will be invoked against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_guard_hook#target_name CloudformationGuardHook#target_name}
	TargetName *string `field:"optional" json:"targetName" yaml:"targetName"`
}

