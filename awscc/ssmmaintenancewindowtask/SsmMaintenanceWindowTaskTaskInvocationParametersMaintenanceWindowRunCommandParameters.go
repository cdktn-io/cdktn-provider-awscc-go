// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskTaskInvocationParametersMaintenanceWindowRunCommandParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#cloudwatch_output_config SsmMaintenanceWindowTask#cloudwatch_output_config}.
	CloudwatchOutputConfig *SsmMaintenanceWindowTaskTaskInvocationParametersMaintenanceWindowRunCommandParametersCloudwatchOutputConfig `field:"optional" json:"cloudwatchOutputConfig" yaml:"cloudwatchOutputConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#comment SsmMaintenanceWindowTask#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#document_hash SsmMaintenanceWindowTask#document_hash}.
	DocumentHash *string `field:"optional" json:"documentHash" yaml:"documentHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#document_hash_type SsmMaintenanceWindowTask#document_hash_type}.
	DocumentHashType *string `field:"optional" json:"documentHashType" yaml:"documentHashType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#document_version SsmMaintenanceWindowTask#document_version}.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#notification_config SsmMaintenanceWindowTask#notification_config}.
	NotificationConfig *SsmMaintenanceWindowTaskTaskInvocationParametersMaintenanceWindowRunCommandParametersNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#output_s3_bucket_name SsmMaintenanceWindowTask#output_s3_bucket_name}.
	OutputS3BucketName *string `field:"optional" json:"outputS3BucketName" yaml:"outputS3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#output_s3_key_prefix SsmMaintenanceWindowTask#output_s3_key_prefix}.
	OutputS3KeyPrefix *string `field:"optional" json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#parameters SsmMaintenanceWindowTask#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#service_role_arn SsmMaintenanceWindowTask#service_role_arn}.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_maintenance_window_task#timeout_seconds SsmMaintenanceWindowTask#timeout_seconds}.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

