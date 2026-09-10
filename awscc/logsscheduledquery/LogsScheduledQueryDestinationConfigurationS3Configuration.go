// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsscheduledquery


type LogsScheduledQueryDestinationConfigurationS3Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/logs_scheduled_query#destination_identifier LogsScheduledQuery#destination_identifier}.
	DestinationIdentifier *string `field:"optional" json:"destinationIdentifier" yaml:"destinationIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/logs_scheduled_query#role_arn LogsScheduledQuery#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

