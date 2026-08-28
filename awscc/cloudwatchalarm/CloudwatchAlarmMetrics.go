// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmMetrics struct {
	// The ID of the account where the metrics are located, if this is a cross-account alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#account_id CloudwatchAlarm#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The math expression to be performed on the returned data, if this object is performing a math expression.
	//
	// This expression can use the ``Id`` of the other metrics to refer to those metrics, and can also use the ``Id`` of other expressions to use the result of those expressions. For more information about metric math expressions, see [Metric Math Syntax and Functions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax) in the *User Guide*.
	//  Within each MetricDataQuery object, you must specify either ``Expression`` or ``MetricStat`` but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#expression CloudwatchAlarm#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// A short name used to tie this object to the results in the response.
	//
	// This name must be unique within a single call to ``GetMetricData``. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#id CloudwatchAlarm#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A human-readable label for this metric or expression.
	//
	// This is especially useful if this is an expression, so that you know what the value represents. If the metric or expression is shown in a CW dashboard widget, the label is shown. If ``Label`` is omitted, CW generates a default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#label CloudwatchAlarm#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The metric to be returned, along with statistics, period, and units.
	//
	// Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.
	//  Within one MetricDataQuery object, you must specify either ``Expression`` or ``MetricStat`` but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#metric_stat CloudwatchAlarm#metric_stat}
	MetricStat *CloudwatchAlarmMetricsMetricStat `field:"optional" json:"metricStat" yaml:"metricStat"`
	// The granularity, in seconds, of the returned data points.
	//
	// For metrics with regular resolution, a period can be as short as one minute (60 seconds) and must be a multiple of 60. For high-resolution metrics that are collected at intervals of less than one minute, the period can be 1, 5, 10, 20, 30, 60, or any multiple of 60. High-resolution metrics are those metrics stored by a ``PutMetricData`` operation that includes a ``StorageResolution of 1 second``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#period CloudwatchAlarm#period}
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// This option indicates whether to return the timestamps and raw data values of this metric.
	//
	// When you create an alarm based on a metric math expression, specify ``True`` for this value for only the one math expression that the alarm is based on. You must specify ``False`` for ``ReturnData`` for all the other metrics and expressions used in the alarm.
	//  This field is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#return_data CloudwatchAlarm#return_data}
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

