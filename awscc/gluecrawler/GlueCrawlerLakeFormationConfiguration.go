// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerLakeFormationConfiguration struct {
	// Required for cross account crawls. For same account crawls as the target data, this can be left as null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_crawler#account_id GlueCrawler#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Specifies whether to use AWS Lake Formation credentials for the crawler instead of the IAM role credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_crawler#use_lake_formation_credentials GlueCrawler#use_lake_formation_credentials}
	UseLakeFormationCredentials interface{} `field:"optional" json:"useLakeFormationCredentials" yaml:"useLakeFormationCredentials"`
}

