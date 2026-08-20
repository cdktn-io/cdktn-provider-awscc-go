// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsscraper


type ApsScraperScraperLoggingConfigurationLoggingDestinationCloudwatchLogs struct {
	// ARN of the CloudWatch log group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/aps_scraper#log_group_arn ApsScraper#log_group_arn}
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
}

