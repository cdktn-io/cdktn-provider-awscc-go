// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfiguration struct {
	// Describes whether to use the default CloudWatch logging configuration for an application.
	//
	// You must set this property to CUSTOM in order to set the LogLevel or MetricsLevel parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisanalyticsv2_application#configuration_type Kinesisanalyticsv2Application#configuration_type}
	ConfigurationType *string `field:"optional" json:"configurationType" yaml:"configurationType"`
	// Describes the verbosity of the CloudWatch Logs for an application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisanalyticsv2_application#log_level Kinesisanalyticsv2Application#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Describes the granularity of the CloudWatch Logs for an application.
	//
	// The Parallelism level is not recommended for applications with a Parallelism over 64 due to excessive costs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisanalyticsv2_application#metrics_level Kinesisanalyticsv2Application#metrics_level}
	MetricsLevel *string `field:"optional" json:"metricsLevel" yaml:"metricsLevel"`
}

