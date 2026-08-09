// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationguardhook


type CloudformationGuardHookTargetFilters struct {
	// List of actions that the hook is going to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudformation_guard_hook#actions CloudformationGuardHook#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// List of invocation points that the hook is going to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudformation_guard_hook#invocation_points CloudformationGuardHook#invocation_points}
	InvocationPoints *[]*string `field:"optional" json:"invocationPoints" yaml:"invocationPoints"`
	// List of type names that the hook is going to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudformation_guard_hook#target_names CloudformationGuardHook#target_names}
	TargetNames *[]*string `field:"optional" json:"targetNames" yaml:"targetNames"`
	// List of hook targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudformation_guard_hook#targets CloudformationGuardHook#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

