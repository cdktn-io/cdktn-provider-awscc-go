// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule


type SagemakerMonitoringScheduleLastMonitoringExecutionSummary struct {
	// The time at which the monitoring job was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#creation_time SagemakerMonitoringSchedule#creation_time}
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// The name of the endpoint used to run the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#endpoint_name SagemakerMonitoringSchedule#endpoint_name}
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Contains the reason a monitoring job failed, if it failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#failure_reason SagemakerMonitoringSchedule#failure_reason}
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// A timestamp that indicates the last time the monitoring job was modified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#last_modified_time SagemakerMonitoringSchedule#last_modified_time}
	LastModifiedTime *string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// The status of the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#monitoring_execution_status SagemakerMonitoringSchedule#monitoring_execution_status}
	MonitoringExecutionStatus *string `field:"optional" json:"monitoringExecutionStatus" yaml:"monitoringExecutionStatus"`
	// The name of the monitoring schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#monitoring_schedule_name SagemakerMonitoringSchedule#monitoring_schedule_name}
	MonitoringScheduleName *string `field:"optional" json:"monitoringScheduleName" yaml:"monitoringScheduleName"`
	// The Amazon Resource Name (ARN) of the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#processing_job_arn SagemakerMonitoringSchedule#processing_job_arn}
	ProcessingJobArn *string `field:"optional" json:"processingJobArn" yaml:"processingJobArn"`
	// The time the monitoring job was scheduled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#scheduled_time SagemakerMonitoringSchedule#scheduled_time}
	ScheduledTime *string `field:"optional" json:"scheduledTime" yaml:"scheduledTime"`
}

