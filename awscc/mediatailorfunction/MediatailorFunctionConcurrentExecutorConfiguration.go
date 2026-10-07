// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction


type MediatailorFunctionConcurrentExecutorConfiguration struct {
	// The list of 1 to 10 child functions that MediaTailor runs in parallel.
	//
	// Each entry specifies a child function to execute and an optional run condition expression that controls whether the function runs. Child functions cannot themselves be executors, and each child function's resolved namespace must be unique across the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#function_list MediatailorFunction#function_list}
	FunctionList interface{} `field:"optional" json:"functionList" yaml:"functionList"`
	// The maximum number of child functions that MediaTailor runs simultaneously.
	//
	// When the list contains more functions than MaxConcurrency, MediaTailor starts additional functions as running ones complete, so that no more than MaxConcurrency functions run at the same time. Valid values are 1 to 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#max_concurrency MediatailorFunction#max_concurrency}
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// A map of output bindings that controls which bindings the executor commits to the session state after all child functions complete.
	//
	// Each key is a namespaced output path, and each value is an expression that MediaTailor evaluates against the combined results of the child functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#output MediatailorFunction#output}
	Output *map[string]*string `field:"optional" json:"output" yaml:"output"`
	// The expression language used to evaluate expressions in the function configuration. Set this to JSONATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#runtime MediatailorFunction#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// The maximum time, in milliseconds, for all child functions to complete.
	//
	// This timeout covers every function in the list, including any HTTP calls the child functions make. If the executor exceeds this timeout, MediaTailor discards all output from the executor and proceeds with default behavior. Valid values are 100 to 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_function#timeout_milliseconds MediatailorFunction#timeout_milliseconds}
	TimeoutMilliseconds *float64 `field:"optional" json:"timeoutMilliseconds" yaml:"timeoutMilliseconds"`
}

