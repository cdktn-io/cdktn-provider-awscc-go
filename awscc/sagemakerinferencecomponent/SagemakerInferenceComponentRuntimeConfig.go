// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentRuntimeConfig struct {
	// The number of copies for the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_inference_component#copy_count SagemakerInferenceComponent#copy_count}
	CopyCount *float64 `field:"optional" json:"copyCount" yaml:"copyCount"`
}

