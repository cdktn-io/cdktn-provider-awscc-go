// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigDataCaptureConfigCaptureOptions struct {
	// Specifies whether the endpoint captures input data or output data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#capture_mode SagemakerEndpointConfigA#capture_mode}
	CaptureMode *string `field:"optional" json:"captureMode" yaml:"captureMode"`
}

