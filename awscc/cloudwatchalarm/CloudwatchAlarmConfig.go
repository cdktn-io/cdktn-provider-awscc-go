// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchAlarmConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#actions_enabled CloudwatchAlarm#actions_enabled}
	ActionsEnabled interface{} `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state.
	//
	// Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutMetricAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html) in the *API Reference*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#alarm_actions CloudwatchAlarm#alarm_actions}
	AlarmActions *[]*string `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// The description of the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#alarm_description CloudwatchAlarm#alarm_description}
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// The name of the alarm.
	//
	// If you don't specify a name, CFN generates a unique physical ID and uses that ID for the alarm name.
	//   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#alarm_name CloudwatchAlarm#alarm_name}
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// The arithmetic operation to use when comparing the specified statistic and threshold.
	//
	// The specified statistic value is used as the first operand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#comparison_operator CloudwatchAlarm#comparison_operator}
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of datapoints that must be breaching to trigger the alarm.
	//
	// This is used only if you are setting an "M out of N" alarm. In that case, this value is the M, and the value that you set for ``EvaluationPeriods`` is the N value. For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
	//  If you omit this parameter, CW uses the same value here that you set for ``EvaluationPeriods``, and the alarm goes to alarm state if that many consecutive periods are breaching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#datapoints_to_alarm CloudwatchAlarm#datapoints_to_alarm}
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// The dimensions for the metric associated with the alarm.
	//
	// For an alarm based on a math expression, you can't specify ``Dimensions``. Instead, you use ``Metrics``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#dimensions CloudwatchAlarm#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Used only for alarms based on percentiles.
	//
	// If ``ignore``, the alarm state does not change during periods with too few data points to be statistically significant. If ``evaluate`` or this parameter is not used, the alarm is always evaluated and possibly changes state no matter how many data points are available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#evaluate_low_sample_count_percentile CloudwatchAlarm#evaluate_low_sample_count_percentile}
	EvaluateLowSampleCountPercentile *string `field:"optional" json:"evaluateLowSampleCountPercentile" yaml:"evaluateLowSampleCountPercentile"`
	// The evaluation criteria for the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#evaluation_criteria CloudwatchAlarm#evaluation_criteria}
	EvaluationCriteria *CloudwatchAlarmEvaluationCriteria `field:"optional" json:"evaluationCriteria" yaml:"evaluationCriteria"`
	// The frequency, in seconds, at which the alarm is evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#evaluation_interval CloudwatchAlarm#evaluation_interval}
	EvaluationInterval *float64 `field:"optional" json:"evaluationInterval" yaml:"evaluationInterval"`
	// The number of periods over which data is compared to the specified threshold.
	//
	// If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an "M out of N" alarm, this value is the N, and ``DatapointsToAlarm`` is the M.
	//  For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#evaluation_periods CloudwatchAlarm#evaluation_periods}
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The evaluation window that the alarm uses to select the range of metric data that it evaluates.
	//
	// This is either a sliding window or a wall clock window. For more information, see [Alarm evaluation windows](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/alarm-evaluation-window.html) in the *CloudWatch User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#evaluation_window CloudwatchAlarm#evaluation_window}
	EvaluationWindow *CloudwatchAlarmEvaluationWindow `field:"optional" json:"evaluationWindow" yaml:"evaluationWindow"`
	// The percentile statistic for the metric associated with the alarm.
	//
	// Specify a value between p0.0 and p100.
	//  For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
	//  For an alarm based on a math expression, you can't specify ``ExtendedStatistic``. Instead, you use ``Metrics``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#extended_statistic CloudwatchAlarm#extended_statistic}
	ExtendedStatistic *string `field:"optional" json:"extendedStatistic" yaml:"extendedStatistic"`
	// The actions to execute when this alarm transitions to the ``INSUFFICIENT_DATA`` state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#insufficient_data_actions CloudwatchAlarm#insufficient_data_actions}
	InsufficientDataActions *[]*string `field:"optional" json:"insufficientDataActions" yaml:"insufficientDataActions"`
	// The name of the metric associated with the alarm.
	//
	// This is required for an alarm based on a metric. For an alarm based on a math expression, you use ``Metrics`` instead and you can't specify ``MetricName``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#metric_name CloudwatchAlarm#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// An array that enables you to create an alarm based on the result of a metric math expression.
	//
	// Each item in the array either retrieves a metric or performs a math expression.
	//  If you specify the ``Metrics`` parameter, you cannot specify ``MetricName``, ``Dimensions``, ``Period``, ``Namespace``, ``Statistic``, ``ExtendedStatistic``, or ``Unit``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#metrics CloudwatchAlarm#metrics}
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// The namespace of the metric associated with the alarm.
	//
	// This is required for an alarm based on a metric. For an alarm based on a math expression, you can't specify ``Namespace`` and you use ``Metrics`` instead.
	//  For a list of namespaces for metrics from AWS services, see [Services That Publish Metrics.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#namespace CloudwatchAlarm#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The actions to execute when this alarm transitions to the ``OK`` state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#ok_actions CloudwatchAlarm#ok_actions}
	OkActions *[]*string `field:"optional" json:"okActions" yaml:"okActions"`
	// The period, in seconds, over which the statistic is applied.
	//
	// This is required for an alarm based on a metric. Valid values are 10, 20, 30, 60, and any multiple of 60.
	//  For an alarm based on a math expression, you can't specify ``Period``, and instead you use the ``Metrics`` parameter.
	//  *Minimum:* 10
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#period CloudwatchAlarm#period}
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The statistic for the metric associated with the alarm, other than percentile.
	//
	// For percentile statistics, use ``ExtendedStatistic``.
	//  For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
	//  For an alarm based on a math expression, you can't specify ``Statistic``. Instead, you use ``Metrics``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#statistic CloudwatchAlarm#statistic}
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// A list of key-value pairs to associate with the alarm.
	//
	// You can associate as many as 50 tags with an alarm. To be able to associate tags with the alarm when you create the alarm, you must have the ``cloudwatch:TagResource`` permission.
	//  Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#tags CloudwatchAlarm#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The value to compare with the specified statistic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#threshold CloudwatchAlarm#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// In an alarm based on an anomaly detection model, this is the ID of the ``ANOMALY_DETECTION_BAND`` function used as the threshold for the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#threshold_metric_id CloudwatchAlarm#threshold_metric_id}
	ThresholdMetricId *string `field:"optional" json:"thresholdMetricId" yaml:"thresholdMetricId"`
	// Sets how this alarm is to handle missing data points.
	//
	// Valid values are ``breaching``, ``notBreaching``, ``ignore``, and ``missing``. For more information, see [Configuring How Alarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the *Amazon User Guide*.
	//  If you omit this parameter, the default behavior of ``missing`` is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#treat_missing_data CloudwatchAlarm#treat_missing_data}
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
	// The unit of the metric associated with the alarm.
	//
	// Specify this only if you are creating an alarm based on a single metric. Do not specify this if you are specifying a ``Metrics`` array.
	//   You can specify the following values: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#unit CloudwatchAlarm#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudwatch_alarm#warm_up_configuration CloudwatchAlarm#warm_up_configuration}.
	WarmUpConfiguration *CloudwatchAlarmWarmUpConfiguration `field:"optional" json:"warmUpConfiguration" yaml:"warmUpConfiguration"`
}

