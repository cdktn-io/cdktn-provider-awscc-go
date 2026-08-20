// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mwaaenvironment


type MwaaEnvironmentLoggingConfiguration struct {
	// Logging configuration for a specific airflow component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mwaa_environment#dag_processing_logs MwaaEnvironment#dag_processing_logs}
	DagProcessingLogs *MwaaEnvironmentLoggingConfigurationDagProcessingLogs `field:"optional" json:"dagProcessingLogs" yaml:"dagProcessingLogs"`
	// Logging configuration for a specific airflow component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mwaa_environment#scheduler_logs MwaaEnvironment#scheduler_logs}
	SchedulerLogs *MwaaEnvironmentLoggingConfigurationSchedulerLogs `field:"optional" json:"schedulerLogs" yaml:"schedulerLogs"`
	// Logging configuration for a specific airflow component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mwaa_environment#task_logs MwaaEnvironment#task_logs}
	TaskLogs *MwaaEnvironmentLoggingConfigurationTaskLogs `field:"optional" json:"taskLogs" yaml:"taskLogs"`
	// Logging configuration for a specific airflow component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mwaa_environment#webserver_logs MwaaEnvironment#webserver_logs}
	WebserverLogs *MwaaEnvironmentLoggingConfigurationWebserverLogs `field:"optional" json:"webserverLogs" yaml:"webserverLogs"`
	// Logging configuration for a specific airflow component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mwaa_environment#worker_logs MwaaEnvironment#worker_logs}
	WorkerLogs *MwaaEnvironmentLoggingConfigurationWorkerLogs `field:"optional" json:"workerLogs" yaml:"workerLogs"`
}

