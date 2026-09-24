// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentSpecificationsContainer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#artifact_url SagemakerInferenceComponent#artifact_url}.
	ArtifactUrl *string `field:"optional" json:"artifactUrl" yaml:"artifactUrl"`
	// The configuration for container metrics scraping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#container_metrics_config SagemakerInferenceComponent#container_metrics_config}
	ContainerMetricsConfig *SagemakerInferenceComponentSpecificationsContainerContainerMetricsConfig `field:"optional" json:"containerMetricsConfig" yaml:"containerMetricsConfig"`
	// Environment variables to specify on the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#environment SagemakerInferenceComponent#environment}
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The image to use for the container that will be materialized for the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#image SagemakerInferenceComponent#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
}

