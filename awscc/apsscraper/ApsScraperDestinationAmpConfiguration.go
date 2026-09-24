// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsscraper


type ApsScraperDestinationAmpConfiguration struct {
	// ARN of an Amazon Managed Prometheus workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_scraper#workspace_arn ApsScraper#workspace_arn}
	WorkspaceArn *string `field:"optional" json:"workspaceArn" yaml:"workspaceArn"`
}

