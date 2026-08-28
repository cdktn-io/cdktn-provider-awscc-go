// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionDurableConfig struct {
	// The maximum time (in seconds) that a durable execution can run before timing out.
	//
	// This timeout applies to the entire durable execution, not individual function invocations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_function#execution_timeout LambdaFunction#execution_timeout}
	ExecutionTimeout *float64 `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// The number of days to retain execution history after a durable execution completes.
	//
	// After this period, execution history is no longer available through the GetDurableExecutionHistory API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_function#retention_period_in_days LambdaFunction#retention_period_in_days}
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
}

