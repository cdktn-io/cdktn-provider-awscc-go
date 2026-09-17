// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigAsyncInferenceConfig struct {
	// Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#client_config SagemakerEndpointConfigA#client_config}
	ClientConfig *SagemakerEndpointConfigAsyncInferenceConfigClientConfig `field:"optional" json:"clientConfig" yaml:"clientConfig"`
	// Specifies the configuration for asynchronous inference invocation outputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#output_config SagemakerEndpointConfigA#output_config}
	OutputConfig *SagemakerEndpointConfigAsyncInferenceConfigOutputConfig `field:"optional" json:"outputConfig" yaml:"outputConfig"`
}

