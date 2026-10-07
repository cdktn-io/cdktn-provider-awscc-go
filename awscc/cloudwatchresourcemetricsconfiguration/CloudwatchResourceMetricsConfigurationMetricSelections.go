// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchresourcemetricsconfiguration


type CloudwatchResourceMetricsConfigurationMetricSelections struct {
	// The list of metric names to include in detailed monitoring for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/cloudwatch_resource_metrics_configuration#include_metrics CloudwatchResourceMetricsConfiguration#include_metrics}
	IncludeMetrics *[]*string `field:"optional" json:"includeMetrics" yaml:"includeMetrics"`
}

