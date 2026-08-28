// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmMetricsMetricStat struct {
	// The metric to return, including the metric name, namespace, and dimensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#metric CloudwatchAlarm#metric}
	Metric *CloudwatchAlarmMetricsMetricStatMetric `field:"optional" json:"metric" yaml:"metric"`
	// The granularity, in seconds, of the returned data points.
	//
	// For metrics with regular resolution, a period can be as short as one minute (60 seconds) and must be a multiple of 60. For high-resolution metrics that are collected at intervals of less than one minute, the period can be 1, 5, 10, 20, 30, 60, or any multiple of 60. High-resolution metrics are those metrics stored by a ``PutMetricData`` call that includes a ``StorageResolution`` of 1 second.
	//  If the ``StartTime`` parameter specifies a time stamp that is greater than 3 hours ago, you must specify the period as follows or no data points in that time range is returned:
	//   +  Start time between 3 hours and 15 days ago - Use a multiple of 60 seconds (1 minute).
	//   +  Start time between 15 and 63 days ago - Use a multiple of 300 seconds (5 minutes).
	//   +  Start time greater than 63 days ago - Use a multiple of 3600 seconds (1 hour).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#period CloudwatchAlarm#period}
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The statistic to return.
	//
	// It can include any CW statistic or extended statistic. For a list of valid values, see the table in [Statistics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Statistic) in the *User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#stat CloudwatchAlarm#stat}
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
	// The unit to use for the returned data points.
	//
	// Valid values are: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#unit CloudwatchAlarm#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

