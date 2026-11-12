// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationRunConfigurationApplicationRestoreConfiguration struct {
	// Specifies how the application should be restored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/kinesisanalyticsv2_application#application_restore_type Kinesisanalyticsv2Application#application_restore_type}
	ApplicationRestoreType *string `field:"optional" json:"applicationRestoreType" yaml:"applicationRestoreType"`
	// The ARN of an existing snapshot to use to restart an application.
	//
	// The application uses this value if RESTORE_FROM_CUSTOM_SNAPSHOT is specified for the ApplicationRestoreType. Specify either SnapshotName or SnapshotArn, not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/kinesisanalyticsv2_application#snapshot_arn Kinesisanalyticsv2Application#snapshot_arn}
	SnapshotArn *string `field:"optional" json:"snapshotArn" yaml:"snapshotArn"`
	// The identifier of an existing snapshot of application state to use to restart an application.
	//
	// The application uses this value if RESTORE_FROM_CUSTOM_SNAPSHOT is specified for the ApplicationRestoreType. Specify either SnapshotName or SnapshotArn, not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/kinesisanalyticsv2_application#snapshot_name Kinesisanalyticsv2Application#snapshot_name}
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
}

