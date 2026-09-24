// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsscraper


type ApsScraperSource struct {
	// Configuration for EKS metrics source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_scraper#eks_configuration ApsScraper#eks_configuration}
	EksConfiguration *ApsScraperSourceEksConfiguration `field:"optional" json:"eksConfiguration" yaml:"eksConfiguration"`
	// Configuration for VPC metrics source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_scraper#vpc_configuration ApsScraper#vpc_configuration}
	VpcConfiguration *ApsScraperSourceVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

