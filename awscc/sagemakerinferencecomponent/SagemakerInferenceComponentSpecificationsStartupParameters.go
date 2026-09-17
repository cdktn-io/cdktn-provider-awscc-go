// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentSpecificationsStartupParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_inference_component#container_startup_health_check_timeout_in_seconds SagemakerInferenceComponent#container_startup_health_check_timeout_in_seconds}.
	ContainerStartupHealthCheckTimeoutInSeconds *float64 `field:"optional" json:"containerStartupHealthCheckTimeoutInSeconds" yaml:"containerStartupHealthCheckTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_inference_component#model_data_download_timeout_in_seconds SagemakerInferenceComponent#model_data_download_timeout_in_seconds}.
	ModelDataDownloadTimeoutInSeconds *float64 `field:"optional" json:"modelDataDownloadTimeoutInSeconds" yaml:"modelDataDownloadTimeoutInSeconds"`
}

