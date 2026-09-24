// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationguardhook


type CloudformationGuardHookStackFiltersStackNames struct {
	// List of stack names that the hook is going to be excluded from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_guard_hook#exclude CloudformationGuardHook#exclude}
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// List of stack names that the hook is going to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudformation_guard_hook#include CloudformationGuardHook#include}
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

