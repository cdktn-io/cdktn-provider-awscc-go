// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsscraper


type ApsScraperRoleConfiguration struct {
	// IAM Role in source account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aps_scraper#source_role_arn ApsScraper#source_role_arn}
	SourceRoleArn *string `field:"optional" json:"sourceRoleArn" yaml:"sourceRoleArn"`
	// IAM Role in the target account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aps_scraper#target_role_arn ApsScraper#target_role_arn}
	TargetRoleArn *string `field:"optional" json:"targetRoleArn" yaml:"targetRoleArn"`
}

