// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinition struct {
	// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#baseline_config SagemakerMonitoringSchedule#baseline_config}
	BaselineConfig *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConfig `field:"optional" json:"baselineConfig" yaml:"baselineConfig"`
	// Sets the environment variables in the Docker container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#environment SagemakerMonitoringSchedule#environment}
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Container image configuration object for the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#monitoring_app_specification SagemakerMonitoringSchedule#monitoring_app_specification}
	MonitoringAppSpecification *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionMonitoringAppSpecification `field:"optional" json:"monitoringAppSpecification" yaml:"monitoringAppSpecification"`
	// The array of inputs for the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#monitoring_inputs SagemakerMonitoringSchedule#monitoring_inputs}
	MonitoringInputs interface{} `field:"optional" json:"monitoringInputs" yaml:"monitoringInputs"`
	// The output configuration for monitoring jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#monitoring_output_config SagemakerMonitoringSchedule#monitoring_output_config}
	MonitoringOutputConfig *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionMonitoringOutputConfig `field:"optional" json:"monitoringOutputConfig" yaml:"monitoringOutputConfig"`
	// Identifies the resources to deploy for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#monitoring_resources SagemakerMonitoringSchedule#monitoring_resources}
	MonitoringResources *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionMonitoringResources `field:"optional" json:"monitoringResources" yaml:"monitoringResources"`
	// Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#network_config SagemakerMonitoringSchedule#network_config}
	NetworkConfig *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#role_arn SagemakerMonitoringSchedule#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Specifies a time limit for how long the monitoring job is allowed to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#stopping_condition SagemakerMonitoringSchedule#stopping_condition}
	StoppingCondition *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionStoppingCondition `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
}

