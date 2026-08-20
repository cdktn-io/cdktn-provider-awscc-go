// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConfigConstraintsResource struct {
	// The Amazon S3 URI for baseline constraint file in Amazon S3 that the current monitoring job should validated against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_monitoring_schedule#s3_uri SagemakerMonitoringSchedule#s3_uri}
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

