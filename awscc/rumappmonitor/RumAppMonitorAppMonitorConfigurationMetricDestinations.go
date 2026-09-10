// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rumappmonitor


type RumAppMonitorAppMonitorConfigurationMetricDestinations struct {
	// Defines the destination to send the metrics to.
	//
	// Valid values are CloudWatch and Evidently. Note: Evidently has been discontinued and is no longer supported - requests with Evidently will be rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#destination RumAppMonitor#destination}
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Evidently has been discontinued and therefore this is no longer an acceptable field.
	//
	// If Destination is CloudWatch, do not use this parameter.
	//
	// This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#destination_arn RumAppMonitor#destination_arn}
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// Evidently has been discontinued and therefore this is no longer an acceptable field.
	//
	// If Destination is CloudWatch, do not use this parameter.
	//
	// This parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#iam_role_arn RumAppMonitor#iam_role_arn}
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// An array of structures which define the metrics that you want to send.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#metric_definitions RumAppMonitor#metric_definitions}
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
}

