// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskTaskInvocationParametersMaintenanceWindowRunCommandParametersCloudwatchOutputConfig struct {
	// The name of the CloudWatch log group where you want to send command output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssm_maintenance_window_task#cloudwatch_log_group_name SsmMaintenanceWindowTask#cloudwatch_log_group_name}
	CloudwatchLogGroupName *string `field:"optional" json:"cloudwatchLogGroupName" yaml:"cloudwatchLogGroupName"`
	// Enables Systems Manager to send command output to CloudWatch Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssm_maintenance_window_task#cloudwatch_output_enabled SsmMaintenanceWindowTask#cloudwatch_output_enabled}
	CloudwatchOutputEnabled interface{} `field:"optional" json:"cloudwatchOutputEnabled" yaml:"cloudwatchOutputEnabled"`
}

