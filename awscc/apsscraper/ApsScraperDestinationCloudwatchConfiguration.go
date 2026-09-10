// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsscraper


type ApsScraperDestinationCloudwatchConfiguration struct {
	// ARN of a CloudWatch dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/aps_scraper#dataset_arn ApsScraper#dataset_arn}
	DatasetArn *string `field:"optional" json:"datasetArn" yaml:"datasetArn"`
}

