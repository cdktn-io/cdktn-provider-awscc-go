// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchmetricstream


type CloudwatchMetricStreamStatisticsConfigurations struct {
	// The additional statistics to stream for the metrics listed in IncludeMetrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudwatch_metric_stream#additional_statistics CloudwatchMetricStream#additional_statistics}
	AdditionalStatistics *[]*string `field:"optional" json:"additionalStatistics" yaml:"additionalStatistics"`
	// An array that defines the metrics that are to have additional statistics streamed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudwatch_metric_stream#include_metrics CloudwatchMetricStream#include_metrics}
	IncludeMetrics interface{} `field:"optional" json:"includeMetrics" yaml:"includeMetrics"`
}

