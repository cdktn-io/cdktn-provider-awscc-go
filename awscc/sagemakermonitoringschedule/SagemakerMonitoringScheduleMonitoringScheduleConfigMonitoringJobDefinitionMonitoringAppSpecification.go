// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionMonitoringAppSpecification struct {
	// An array of arguments for the container used to run the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#container_arguments SagemakerMonitoringSchedule#container_arguments}
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Specifies the entrypoint for a container used to run the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#container_entrypoint SagemakerMonitoringSchedule#container_entrypoint}
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// The container image to be run by the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#image_uri SagemakerMonitoringSchedule#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// An Amazon S3 URI to a script that is called after analysis has been performed.
	//
	// Applicable only for the built-in (first party) containers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#post_analytics_processor_source_uri SagemakerMonitoringSchedule#post_analytics_processor_source_uri}
	PostAnalyticsProcessorSourceUri *string `field:"optional" json:"postAnalyticsProcessorSourceUri" yaml:"postAnalyticsProcessorSourceUri"`
	// An Amazon S3 URI to a script that is called per row prior to running analysis.
	//
	// It can base64 decode the payload and convert it into a flatted json so that the built-in container can use the converted data. Applicable only for the built-in (first party) containers
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_monitoring_schedule#record_preprocessor_source_uri SagemakerMonitoringSchedule#record_preprocessor_source_uri}
	RecordPreprocessorSourceUri *string `field:"optional" json:"recordPreprocessorSourceUri" yaml:"recordPreprocessorSourceUri"`
}

