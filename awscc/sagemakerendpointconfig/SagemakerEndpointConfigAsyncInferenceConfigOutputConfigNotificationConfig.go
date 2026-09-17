// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigAsyncInferenceConfigOutputConfigNotificationConfig struct {
	// Amazon SNS topic to post a notification to when an inference fails.
	//
	// If no topic is provided, no notification is sent on failure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#error_topic SagemakerEndpointConfigA#error_topic}
	ErrorTopic *string `field:"optional" json:"errorTopic" yaml:"errorTopic"`
	// The Amazon SNS topics where you want the inference response to be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#include_inference_response_in SagemakerEndpointConfigA#include_inference_response_in}
	IncludeInferenceResponseIn *[]*string `field:"optional" json:"includeInferenceResponseIn" yaml:"includeInferenceResponseIn"`
	// Amazon SNS topic to post a notification to when an inference completes successfully.
	//
	// If no topic is provided, no notification is sent on success.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#success_topic SagemakerEndpointConfigA#success_topic}
	SuccessTopic *string `field:"optional" json:"successTopic" yaml:"successTopic"`
}

