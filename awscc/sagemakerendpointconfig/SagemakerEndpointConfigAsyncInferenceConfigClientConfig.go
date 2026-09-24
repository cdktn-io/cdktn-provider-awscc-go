// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigAsyncInferenceConfigClientConfig struct {
	// The maximum number of concurrent requests sent by the SageMaker client to the model container.
	//
	// If no value is provided, SageMaker will choose an optimal value for you.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#max_concurrent_invocations_per_instance SagemakerEndpointConfigA#max_concurrent_invocations_per_instance}
	MaxConcurrentInvocationsPerInstance *float64 `field:"optional" json:"maxConcurrentInvocationsPerInstance" yaml:"maxConcurrentInvocationsPerInstance"`
}

