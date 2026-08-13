// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsscraper

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApsScraperConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Scraper metrics destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/aps_scraper#destination ApsScraper#destination}
	Destination *ApsScraperDestination `field:"required" json:"destination" yaml:"destination"`
	// Scraper configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/aps_scraper#scrape_configuration ApsScraper#scrape_configuration}
	ScrapeConfiguration *ApsScraperScrapeConfiguration `field:"required" json:"scrapeConfiguration" yaml:"scrapeConfiguration"`
	// Scraper metrics source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/aps_scraper#source ApsScraper#source}
	Source *ApsScraperSource `field:"required" json:"source" yaml:"source"`
	// Scraper alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/aps_scraper#alias ApsScraper#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Role configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/aps_scraper#role_configuration ApsScraper#role_configuration}
	RoleConfiguration *ApsScraperRoleConfiguration `field:"optional" json:"roleConfiguration" yaml:"roleConfiguration"`
	// Configuration for scraper logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/aps_scraper#scraper_logging_configuration ApsScraper#scraper_logging_configuration}
	ScraperLoggingConfiguration *ApsScraperScraperLoggingConfiguration `field:"optional" json:"scraperLoggingConfiguration" yaml:"scraperLoggingConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/aps_scraper#tags ApsScraper#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

