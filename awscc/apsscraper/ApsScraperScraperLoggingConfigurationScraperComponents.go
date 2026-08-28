// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsscraper


type ApsScraperScraperLoggingConfigurationScraperComponents struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aps_scraper#config ApsScraper#config}.
	Config *ApsScraperScraperLoggingConfigurationScraperComponentsConfig `field:"optional" json:"config" yaml:"config"`
	// Type of scraper component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aps_scraper#type ApsScraper#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

