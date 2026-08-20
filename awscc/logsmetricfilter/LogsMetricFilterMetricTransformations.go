// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsmetricfilter


type LogsMetricFilterMetricTransformations struct {
	// The name of the CloudWatch metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#metric_name LogsMetricFilter#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// A custom namespace to contain your metric in CloudWatch.
	//
	// Use namespaces to group together metrics that are similar. For more information, see [Namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Namespace).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#metric_namespace LogsMetricFilter#metric_namespace}
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The value that is published to the CloudWatch metric.
	//
	// For example, if you're counting the occurrences of a particular term like ``Error``, specify 1 for the metric value. If you're counting the number of bytes transferred, reference the value that is in the log event by using $. followed by the name of the field that you specified in the filter pattern, such as ``$.size``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#metric_value LogsMetricFilter#metric_value}
	MetricValue *string `field:"required" json:"metricValue" yaml:"metricValue"`
	// (Optional) The value to emit when a filter pattern does not match a log event.
	//
	// This value can be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#default_value LogsMetricFilter#default_value}
	DefaultValue *float64 `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The fields to use as dimensions for the metric.
	//
	// One metric filter can include as many as three dimensions.
	//   Metrics extracted from log events are charged as custom metrics. To prevent unexpected high charges, do not specify high-cardinality fields such as ``IPAddress`` or ``requestID`` as dimensions. Each different value found for a dimension is treated as a separate metric and accrues charges as a separate custom metric.
	//  CloudWatch Logs disables a metric filter if it generates 1000 different name/value pairs for your specified dimensions within a certain amount of time. This helps to prevent accidental high charges.
	//  You can also set up a billing alarm to alert you if your charges are higher than expected. For more information, see [Creating a Billing Alarm to Monitor Your Estimated Charges](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/monitor_estimated_charges_with_cloudwatch.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#dimensions LogsMetricFilter#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The unit to assign to the metric. If you omit this, the unit is set as ``None``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_metric_filter#unit LogsMetricFilter#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

