// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionCapacityProviderConfigLambdaManagedInstancesCapacityProviderConfig struct {
	// The Amazon Resource Name (ARN) of the capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_function#capacity_provider_arn LambdaFunction#capacity_provider_arn}
	CapacityProviderArn *string `field:"optional" json:"capacityProviderArn" yaml:"capacityProviderArn"`
	// The amount of memory in GiB allocated per vCPU for execution environments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_function#execution_environment_memory_gi_b_per_v_cpu LambdaFunction#execution_environment_memory_gi_b_per_v_cpu}
	ExecutionEnvironmentMemoryGiBPerVCpu *float64 `field:"optional" json:"executionEnvironmentMemoryGiBPerVCpu" yaml:"executionEnvironmentMemoryGiBPerVCpu"`
	// The maximum number of concurrent executions that can run on each execution environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_function#per_execution_environment_max_concurrency LambdaFunction#per_execution_environment_max_concurrency}
	PerExecutionEnvironmentMaxConcurrency *float64 `field:"optional" json:"perExecutionEnvironmentMaxConcurrency" yaml:"perExecutionEnvironmentMaxConcurrency"`
}

