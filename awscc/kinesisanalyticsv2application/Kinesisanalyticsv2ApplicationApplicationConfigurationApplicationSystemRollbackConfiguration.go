// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSystemRollbackConfiguration struct {
	// Describes whether system initiated rollbacks are enabled for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisanalyticsv2_application#rollback_enabled Kinesisanalyticsv2Application#rollback_enabled}
	RollbackEnabled interface{} `field:"optional" json:"rollbackEnabled" yaml:"rollbackEnabled"`
}

