// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationlambdahook


type CloudformationLambdaHookStackFilters struct {
	// Attribute to specify the filtering behavior.
	//
	// ANY will make the Hook pass if one filter matches. ALL will make the Hook pass if all filters match
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_lambda_hook#filtering_criteria CloudformationLambdaHook#filtering_criteria}
	FilteringCriteria *string `field:"optional" json:"filteringCriteria" yaml:"filteringCriteria"`
	// List of stack names as filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_lambda_hook#stack_names CloudformationLambdaHook#stack_names}
	StackNames *CloudformationLambdaHookStackFiltersStackNames `field:"optional" json:"stackNames" yaml:"stackNames"`
	// List of stack roles that are performing the stack operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_lambda_hook#stack_roles CloudformationLambdaHook#stack_roles}
	StackRoles *CloudformationLambdaHookStackFiltersStackRoles `field:"optional" json:"stackRoles" yaml:"stackRoles"`
}

