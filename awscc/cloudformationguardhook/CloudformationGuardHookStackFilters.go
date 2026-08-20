// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationguardhook


type CloudformationGuardHookStackFilters struct {
	// Attribute to specify the filtering behavior.
	//
	// ANY will make the Hook pass if one filter matches. ALL will make the Hook pass if all filters match
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_guard_hook#filtering_criteria CloudformationGuardHook#filtering_criteria}
	FilteringCriteria *string `field:"optional" json:"filteringCriteria" yaml:"filteringCriteria"`
	// List of stack names as filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_guard_hook#stack_names CloudformationGuardHook#stack_names}
	StackNames *CloudformationGuardHookStackFiltersStackNames `field:"optional" json:"stackNames" yaml:"stackNames"`
	// List of stack roles that are performing the stack operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_guard_hook#stack_roles CloudformationGuardHook#stack_roles}
	StackRoles *CloudformationGuardHookStackFiltersStackRoles `field:"optional" json:"stackRoles" yaml:"stackRoles"`
}

