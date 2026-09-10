// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsscraper


type ApsScraperScraperLoggingConfigurationLoggingDestination struct {
	// Represents a cloudwatch logs destination for scraper logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/aps_scraper#cloudwatch_logs ApsScraper#cloudwatch_logs}
	CloudwatchLogs *ApsScraperScraperLoggingConfigurationLoggingDestinationCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
}

