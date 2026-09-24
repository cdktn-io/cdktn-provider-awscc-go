// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsscheduledquery


type LogsScheduledQueryDestinationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/logs_scheduled_query#s3_configuration LogsScheduledQuery#s3_configuration}.
	S3Configuration *LogsScheduledQueryDestinationConfigurationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

