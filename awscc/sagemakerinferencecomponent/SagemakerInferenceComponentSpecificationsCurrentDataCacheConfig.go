// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentSpecificationsCurrentDataCacheConfig struct {
	// Whether the endpoint caches the model artifacts and container image on each instance it provisions for the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_inference_component#enable_caching SagemakerInferenceComponent#enable_caching}
	EnableCaching interface{} `field:"optional" json:"enableCaching" yaml:"enableCaching"`
}

