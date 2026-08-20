// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectmetric


type ConnectMetricMetricCalculationCalculationComponents struct {
	// Metric calculation component alias for use within a calculation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_metric#alias ConnectMetric#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_metric#metric_filters ConnectMetric#metric_filters}.
	MetricFilters interface{} `field:"optional" json:"metricFilters" yaml:"metricFilters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_metric#metric_id ConnectMetric#metric_id}.
	MetricId *string `field:"optional" json:"metricId" yaml:"metricId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_metric#metric_name ConnectMetric#metric_name}.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
}

