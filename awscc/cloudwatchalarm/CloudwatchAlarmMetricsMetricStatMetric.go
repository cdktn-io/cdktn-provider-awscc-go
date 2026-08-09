// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmMetricsMetricStatMetric struct {
	// The metric dimensions that you want to be used for the metric that the alarm will watch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudwatch_alarm#dimensions CloudwatchAlarm#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The name of the metric that you want the alarm to watch. This is a required field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudwatch_alarm#metric_name CloudwatchAlarm#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The namespace of the metric that the alarm will watch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudwatch_alarm#namespace CloudwatchAlarm#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

