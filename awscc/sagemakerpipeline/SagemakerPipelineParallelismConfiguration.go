// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerpipeline


type SagemakerPipelineParallelismConfiguration struct {
	// Maximum parallel execution steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_pipeline#max_parallel_execution_steps SagemakerPipeline#max_parallel_execution_steps}
	MaxParallelExecutionSteps *float64 `field:"optional" json:"maxParallelExecutionSteps" yaml:"maxParallelExecutionSteps"`
}

