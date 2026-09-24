// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigProductionVariantsServerlessConfig struct {
	// The maximum number of concurrent invocations your serverless endpoint can process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#max_concurrency SagemakerEndpointConfigA#max_concurrency}
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The memory size of your serverless endpoint.
	//
	// Valid values are in 1 GB increments: 1024 MB, 2048 MB, 3072 MB, 4096 MB, 5120 MB, or 6144 MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#memory_size_in_mb SagemakerEndpointConfigA#memory_size_in_mb}
	MemorySizeInMb *float64 `field:"optional" json:"memorySizeInMb" yaml:"memorySizeInMb"`
	// The amount of provisioned concurrency to allocate for the serverless endpoint. Should be less than or equal to MaxConcurrency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#provisioned_concurrency SagemakerEndpointConfigA#provisioned_concurrency}
	ProvisionedConcurrency *float64 `field:"optional" json:"provisionedConcurrency" yaml:"provisionedConcurrency"`
}

