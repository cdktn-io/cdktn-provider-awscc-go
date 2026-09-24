// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigMetricsConfig struct {
	// Specifies whether to enable detailed observability for the endpoint.
	//
	// When set to true, the endpoint publishes container-level inference metrics, per-GPU metrics, per-instance host metrics, and inference component placement metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#enable_detailed_observability SagemakerEndpointConfigA#enable_detailed_observability}
	EnableDetailedObservability interface{} `field:"optional" json:"enableDetailedObservability" yaml:"enableDetailedObservability"`
	// Specifies whether to enable enhanced metrics for the endpoint.
	//
	// Enhanced metrics provide utilization and invocation data at instance and container granularity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#enable_enhanced_metrics SagemakerEndpointConfigA#enable_enhanced_metrics}
	EnableEnhancedMetrics interface{} `field:"optional" json:"enableEnhancedMetrics" yaml:"enableEnhancedMetrics"`
	// The interval, in seconds, at which the endpoint publishes metrics to Amazon CloudWatch.
	//
	// Valid values are 10, 30, 60, 120, 180, 240, and 300. The default is 60.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#metric_publish_frequency_in_seconds SagemakerEndpointConfigA#metric_publish_frequency_in_seconds}
	MetricPublishFrequencyInSeconds *float64 `field:"optional" json:"metricPublishFrequencyInSeconds" yaml:"metricPublishFrequencyInSeconds"`
}

