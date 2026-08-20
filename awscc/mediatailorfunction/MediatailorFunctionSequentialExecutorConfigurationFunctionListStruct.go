// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction


type MediatailorFunctionSequentialExecutorConfigurationFunctionListStruct struct {
	// The identifier of the function to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediatailor_function#function_id MediatailorFunction#function_id}
	FunctionId *string `field:"optional" json:"functionId" yaml:"functionId"`
	// A conditional expression that determines whether this function should execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediatailor_function#run_condition MediatailorFunction#run_condition}
	RunCondition *string `field:"optional" json:"runCondition" yaml:"runCondition"`
}

