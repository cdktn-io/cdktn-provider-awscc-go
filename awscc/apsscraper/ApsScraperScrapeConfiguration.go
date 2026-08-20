// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsscraper


type ApsScraperScrapeConfiguration struct {
	// Prometheus compatible scrape configuration in base64 encoded blob format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/aps_scraper#configuration_blob ApsScraper#configuration_blob}
	ConfigurationBlob *string `field:"optional" json:"configurationBlob" yaml:"configurationBlob"`
}

