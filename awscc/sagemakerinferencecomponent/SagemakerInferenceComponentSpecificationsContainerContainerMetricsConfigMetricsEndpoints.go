// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentSpecificationsContainerContainerMetricsConfigMetricsEndpoints struct {
	// The interval, in seconds, at which container metrics scraped from the endpoint are published to Amazon CloudWatch.
	//
	// Valid values per the SageMaker API Reference are 10, 30, 60, 120, 180, 240 and 300; the service validates the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#metric_publish_frequency_in_seconds SagemakerInferenceComponent#metric_publish_frequency_in_seconds}
	MetricPublishFrequencyInSeconds *float64 `field:"optional" json:"metricPublishFrequencyInSeconds" yaml:"metricPublishFrequencyInSeconds"`
	// The path to the Prometheus formatted metrics endpoint exposed by the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#metrics_endpoint_path SagemakerInferenceComponent#metrics_endpoint_path}
	MetricsEndpointPath *string `field:"optional" json:"metricsEndpointPath" yaml:"metricsEndpointPath"`
}

