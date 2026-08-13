// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsscraper


type ApsScraperScraperLoggingConfiguration struct {
	// Destination for scraper logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/aps_scraper#logging_destination ApsScraper#logging_destination}
	LoggingDestination *ApsScraperScraperLoggingConfigurationLoggingDestination `field:"optional" json:"loggingDestination" yaml:"loggingDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/aps_scraper#scraper_components ApsScraper#scraper_components}.
	ScraperComponents interface{} `field:"optional" json:"scraperComponents" yaml:"scraperComponents"`
}

