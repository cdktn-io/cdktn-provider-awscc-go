// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationlambdahook


type CloudformationLambdaHookStackFiltersStackNames struct {
	// List of stack names that the hook is going to be excluded from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_lambda_hook#exclude CloudformationLambdaHook#exclude}
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// List of stack names that the hook is going to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_lambda_hook#include CloudformationLambdaHook#include}
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

