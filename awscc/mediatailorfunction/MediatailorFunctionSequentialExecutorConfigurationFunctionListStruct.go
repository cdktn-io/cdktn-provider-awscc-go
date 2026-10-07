// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction


type MediatailorFunctionSequentialExecutorConfigurationFunctionListStruct struct {
	// An optional alternate name for the child function within the executor.
	//
	// MediaTailor uses this value as the namespace for the child function's output. If omitted, MediaTailor uses the function identifier. The resolved namespace must be unique across all child functions in the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#alias MediatailorFunction#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The identifier of the child function to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#function_id MediatailorFunction#function_id}
	FunctionId *string `field:"optional" json:"functionId" yaml:"functionId"`
	// An optional expression that evaluates to a boolean.
	//
	// MediaTailor evaluates this expression immediately before running the child function, using the accumulated state at that point. If the expression evaluates to false, MediaTailor skips the child function. If omitted, the child function always runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#run_condition MediatailorFunction#run_condition}
	RunCondition *string `field:"optional" json:"runCondition" yaml:"runCondition"`
}

