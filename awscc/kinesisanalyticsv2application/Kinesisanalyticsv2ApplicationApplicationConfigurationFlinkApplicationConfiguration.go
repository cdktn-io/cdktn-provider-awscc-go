// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration struct {
	// Describes an application's checkpointing configuration.
	//
	// Checkpointing is the process of persisting application state for fault tolerance. For more information, see Checkpoints for Fault Tolerance in the Apache Flink Documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisanalyticsv2_application#checkpoint_configuration Kinesisanalyticsv2Application#checkpoint_configuration}
	CheckpointConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfiguration `field:"optional" json:"checkpointConfiguration" yaml:"checkpointConfiguration"`
	// Describes configuration parameters for Amazon CloudWatch logging for an application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisanalyticsv2_application#monitoring_configuration Kinesisanalyticsv2Application#monitoring_configuration}
	MonitoringConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfiguration `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// Describes parameters for how an application executes multiple tasks simultaneously.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisanalyticsv2_application#parallelism_configuration Kinesisanalyticsv2Application#parallelism_configuration}
	ParallelismConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfiguration `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
}

