// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskTaskInvocationParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#maintenance_window_automation_parameters SsmMaintenanceWindowTask#maintenance_window_automation_parameters}.
	MaintenanceWindowAutomationParameters *SsmMaintenanceWindowTaskTaskInvocationParametersMaintenanceWindowAutomationParameters `field:"optional" json:"maintenanceWindowAutomationParameters" yaml:"maintenanceWindowAutomationParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#maintenance_window_lambda_parameters SsmMaintenanceWindowTask#maintenance_window_lambda_parameters}.
	MaintenanceWindowLambdaParameters *SsmMaintenanceWindowTaskTaskInvocationParametersMaintenanceWindowLambdaParameters `field:"optional" json:"maintenanceWindowLambdaParameters" yaml:"maintenanceWindowLambdaParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#maintenance_window_run_command_parameters SsmMaintenanceWindowTask#maintenance_window_run_command_parameters}.
	MaintenanceWindowRunCommandParameters *SsmMaintenanceWindowTaskTaskInvocationParametersMaintenanceWindowRunCommandParameters `field:"optional" json:"maintenanceWindowRunCommandParameters" yaml:"maintenanceWindowRunCommandParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#maintenance_window_step_functions_parameters SsmMaintenanceWindowTask#maintenance_window_step_functions_parameters}.
	MaintenanceWindowStepFunctionsParameters *SsmMaintenanceWindowTaskTaskInvocationParametersMaintenanceWindowStepFunctionsParameters `field:"optional" json:"maintenanceWindowStepFunctionsParameters" yaml:"maintenanceWindowStepFunctionsParameters"`
}

