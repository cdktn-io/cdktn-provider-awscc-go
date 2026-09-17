// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluedataqualityruleset


type GlueDataQualityRulesetTargetTable struct {
	// The name of the database where the AWS Glue table exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_data_quality_ruleset#database_name GlueDataQualityRuleset#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the AWS Glue table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_data_quality_ruleset#table_name GlueDataQualityRuleset#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

