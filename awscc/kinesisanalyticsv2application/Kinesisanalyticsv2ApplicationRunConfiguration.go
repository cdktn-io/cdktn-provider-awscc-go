// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationRunConfiguration struct {
	// Describes the restore behavior of a restarting application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#application_restore_configuration Kinesisanalyticsv2Application#application_restore_configuration}
	ApplicationRestoreConfiguration *Kinesisanalyticsv2ApplicationRunConfigurationApplicationRestoreConfiguration `field:"optional" json:"applicationRestoreConfiguration" yaml:"applicationRestoreConfiguration"`
	// Describes the starting parameters for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#flink_run_configuration Kinesisanalyticsv2Application#flink_run_configuration}
	FlinkRunConfiguration *Kinesisanalyticsv2ApplicationRunConfigurationFlinkRunConfiguration `field:"optional" json:"flinkRunConfiguration" yaml:"flinkRunConfiguration"`
}

