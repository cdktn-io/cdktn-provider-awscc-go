// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizemetricattribution


type PersonalizeMetricAttributionMetrics struct {
	// The metric's event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/personalize_metric_attribution#event_type PersonalizeMetricAttribution#event_type}
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// The attribute's expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/personalize_metric_attribution#expression PersonalizeMetricAttribution#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The metric's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/personalize_metric_attribution#metric_name PersonalizeMetricAttribution#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
}

