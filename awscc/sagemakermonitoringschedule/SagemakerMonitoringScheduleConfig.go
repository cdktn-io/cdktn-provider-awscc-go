// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerMonitoringScheduleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The configuration object that specifies the monitoring schedule and defines the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_monitoring_schedule#monitoring_schedule_config SagemakerMonitoringSchedule#monitoring_schedule_config}
	MonitoringScheduleConfig *SagemakerMonitoringScheduleMonitoringScheduleConfig `field:"required" json:"monitoringScheduleConfig" yaml:"monitoringScheduleConfig"`
	// The name of the monitoring schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_monitoring_schedule#monitoring_schedule_name SagemakerMonitoringSchedule#monitoring_schedule_name}
	MonitoringScheduleName *string `field:"required" json:"monitoringScheduleName" yaml:"monitoringScheduleName"`
	// The name of the endpoint used to run the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_monitoring_schedule#endpoint_name SagemakerMonitoringSchedule#endpoint_name}
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Contains the reason a monitoring job failed, if it failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_monitoring_schedule#failure_reason SagemakerMonitoringSchedule#failure_reason}
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// Describes metadata on the last execution to run, if there was one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_monitoring_schedule#last_monitoring_execution_summary SagemakerMonitoringSchedule#last_monitoring_execution_summary}
	LastMonitoringExecutionSummary *SagemakerMonitoringScheduleLastMonitoringExecutionSummary `field:"optional" json:"lastMonitoringExecutionSummary" yaml:"lastMonitoringExecutionSummary"`
	// The status of a schedule job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_monitoring_schedule#monitoring_schedule_status SagemakerMonitoringSchedule#monitoring_schedule_status}
	MonitoringScheduleStatus *string `field:"optional" json:"monitoringScheduleStatus" yaml:"monitoringScheduleStatus"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_monitoring_schedule#tags SagemakerMonitoringSchedule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

