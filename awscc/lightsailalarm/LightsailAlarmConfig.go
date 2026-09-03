// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailalarm

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailAlarmConfig struct {
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
	// The name for the alarm.
	//
	// Specify the name of an existing alarm to update, and overwrite the previous configuration of the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_alarm#alarm_name LightsailAlarm#alarm_name}
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// The arithmetic operation to use when comparing the specified statistic to the threshold.
	//
	// The specified statistic value is used as the first operand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_alarm#comparison_operator LightsailAlarm#comparison_operator}
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of most recent periods over which data is compared to the specified threshold.
	//
	// If you are setting an "M out of N" alarm, this value (evaluationPeriods) is the N.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_alarm#evaluation_periods LightsailAlarm#evaluation_periods}
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The name of the metric to associate with the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_alarm#metric_name LightsailAlarm#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The name of the Lightsail resource that the alarm monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_alarm#monitored_resource_name LightsailAlarm#monitored_resource_name}
	MonitoredResourceName *string `field:"required" json:"monitoredResourceName" yaml:"monitoredResourceName"`
	// The value against which the specified statistic is compared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_alarm#threshold LightsailAlarm#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// The contact protocols to use for the alarm, such as Email, SMS (text messaging), or both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_alarm#contact_protocols LightsailAlarm#contact_protocols}
	ContactProtocols *[]*string `field:"optional" json:"contactProtocols" yaml:"contactProtocols"`
	// The number of data points that must be not within the specified threshold to trigger the alarm.
	//
	// If you are setting an "M out of N" alarm, this value (datapointsToAlarm) is the M.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_alarm#datapoints_to_alarm LightsailAlarm#datapoints_to_alarm}
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// Indicates whether the alarm is enabled. Notifications are enabled by default if you don't specify this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_alarm#notification_enabled LightsailAlarm#notification_enabled}
	NotificationEnabled interface{} `field:"optional" json:"notificationEnabled" yaml:"notificationEnabled"`
	// The alarm states that trigger a notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_alarm#notification_triggers LightsailAlarm#notification_triggers}
	NotificationTriggers *[]*string `field:"optional" json:"notificationTriggers" yaml:"notificationTriggers"`
	// Sets how this alarm will handle missing data points.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_alarm#treat_missing_data LightsailAlarm#treat_missing_data}
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
}

