// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchmetricstream


type CloudwatchMetricStreamExcludeFilters struct {
	// Only metrics with MetricNames matching these values will be streamed. Must be set together with Namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudwatch_metric_stream#metric_names CloudwatchMetricStream#metric_names}
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
	// Only metrics with Namespace matching this value will be streamed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudwatch_metric_stream#namespace CloudwatchMetricStream#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

