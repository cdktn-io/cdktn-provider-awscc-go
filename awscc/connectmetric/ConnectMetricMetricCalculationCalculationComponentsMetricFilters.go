// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectmetric


type ConnectMetricMetricCalculationCalculationComponentsMetricFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_metric#boolean_condition ConnectMetric#boolean_condition}.
	BooleanCondition *ConnectMetricMetricCalculationCalculationComponentsMetricFiltersBooleanCondition `field:"optional" json:"booleanCondition" yaml:"booleanCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_metric#metric_filter_key ConnectMetric#metric_filter_key}.
	MetricFilterKey *string `field:"optional" json:"metricFilterKey" yaml:"metricFilterKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_metric#negate ConnectMetric#negate}.
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_metric#number_condition ConnectMetric#number_condition}.
	NumberCondition *ConnectMetricMetricCalculationCalculationComponentsMetricFiltersNumberCondition `field:"optional" json:"numberCondition" yaml:"numberCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_metric#string_condition ConnectMetric#string_condition}.
	StringCondition *ConnectMetricMetricCalculationCalculationComponentsMetricFiltersStringCondition `field:"optional" json:"stringCondition" yaml:"stringCondition"`
}

