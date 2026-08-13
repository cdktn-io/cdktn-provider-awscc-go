// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecurityprofile


type IotSecurityProfileAdditionalMetricsToRetainV2 struct {
	// Flag to enable/disable metrics export for metric to be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#export_metric IotSecurityProfile#export_metric}
	ExportMetric interface{} `field:"optional" json:"exportMetric" yaml:"exportMetric"`
	// What is measured by the behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#metric IotSecurityProfile#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// The dimension of a metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#metric_dimension IotSecurityProfile#metric_dimension}
	MetricDimension *IotSecurityProfileAdditionalMetricsToRetainV2MetricDimension `field:"optional" json:"metricDimension" yaml:"metricDimension"`
}

