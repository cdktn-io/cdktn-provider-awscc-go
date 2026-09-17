// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskLoggingInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#region SsmMaintenanceWindowTask#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#s3_bucket SsmMaintenanceWindowTask#s3_bucket}.
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#s3_prefix SsmMaintenanceWindowTask#s3_prefix}.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
}

