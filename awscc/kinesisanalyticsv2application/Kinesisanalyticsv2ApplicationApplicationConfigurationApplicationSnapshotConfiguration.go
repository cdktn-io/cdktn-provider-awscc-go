// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfiguration struct {
	// The snapshot mode of a Flink-based Kinesis Data Analytics application.
	//
	// Changing from APPLICATION_SNAPSHOTS_LEGACY to ENHANCED_SNAPSHOTS is one-way; an application in ENHANCED_SNAPSHOTS mode cannot return to APPLICATION_SNAPSHOTS_LEGACY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/kinesisanalyticsv2_application#snapshot_mode Kinesisanalyticsv2Application#snapshot_mode}
	SnapshotMode *string `field:"optional" json:"snapshotMode" yaml:"snapshotMode"`
	// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/kinesisanalyticsv2_application#snapshots_enabled Kinesisanalyticsv2Application#snapshots_enabled}
	SnapshotsEnabled interface{} `field:"optional" json:"snapshotsEnabled" yaml:"snapshotsEnabled"`
}

