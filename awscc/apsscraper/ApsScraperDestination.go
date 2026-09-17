// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsscraper


type ApsScraperDestination struct {
	// Configuration for Amazon Managed Prometheus metrics destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/aps_scraper#amp_configuration ApsScraper#amp_configuration}
	AmpConfiguration *ApsScraperDestinationAmpConfiguration `field:"optional" json:"ampConfiguration" yaml:"ampConfiguration"`
	// Configuration for CloudWatch metrics destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/aps_scraper#cloudwatch_configuration ApsScraper#cloudwatch_configuration}
	CloudwatchConfiguration *ApsScraperDestinationCloudwatchConfiguration `field:"optional" json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
}

