// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecurityprofile


type IotSecurityProfileBehaviors struct {
	// The criteria by which the behavior is determined to be normal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_security_profile#criteria IotSecurityProfile#criteria}
	Criteria *IotSecurityProfileBehaviorsCriteria `field:"optional" json:"criteria" yaml:"criteria"`
	// Flag to enable/disable metrics export for metric to be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_security_profile#export_metric IotSecurityProfile#export_metric}
	ExportMetric interface{} `field:"optional" json:"exportMetric" yaml:"exportMetric"`
	// What is measured by the behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_security_profile#metric IotSecurityProfile#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// The dimension of a metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_security_profile#metric_dimension IotSecurityProfile#metric_dimension}
	MetricDimension *IotSecurityProfileBehaviorsMetricDimension `field:"optional" json:"metricDimension" yaml:"metricDimension"`
	// The name for the behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_security_profile#name IotSecurityProfile#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Manage Detect alarm SNS notifications by setting behavior notification to on or suppressed.
	//
	// Detect will continue to performing device behavior evaluations. However, suppressed alarms wouldn't be forwarded for SNS notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_security_profile#suppress_alerts IotSecurityProfile#suppress_alerts}
	SuppressAlerts interface{} `field:"optional" json:"suppressAlerts" yaml:"suppressAlerts"`
}

