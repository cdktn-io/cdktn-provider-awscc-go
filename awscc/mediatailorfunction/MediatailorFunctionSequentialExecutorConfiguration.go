// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction


type MediatailorFunctionSequentialExecutorConfiguration struct {
	// An ordered list of 1 to 10 steps.
	//
	// Each step specifies a child function to execute and an optional run condition expression that controls whether the step runs. MediaTailor executes the steps in order, passing data between steps through temporary data. Each step's resolved namespace must be unique across the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#function_list MediatailorFunction#function_list}
	FunctionList interface{} `field:"optional" json:"functionList" yaml:"functionList"`
	// A map of output bindings that controls which bindings the sequence commits to the session state after all steps complete.
	//
	// Each key is a namespaced output path, and each value is an expression that MediaTailor evaluates against the accumulated results of the steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#output MediatailorFunction#output}
	Output *map[string]*string `field:"optional" json:"output" yaml:"output"`
	// The expression language used to evaluate expressions in the function configuration. Set this to JSONATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#runtime MediatailorFunction#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// The maximum time, in milliseconds, for the entire sequence to complete.
	//
	// This timeout covers all steps, including any HTTP calls made by child functions. If the sequence exceeds this timeout, MediaTailor discards all output from the sequence and proceeds with default behavior. Valid values are 100 to 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#timeout_milliseconds MediatailorFunction#timeout_milliseconds}
	TimeoutMilliseconds *float64 `field:"optional" json:"timeoutMilliseconds" yaml:"timeoutMilliseconds"`
}

