// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowMetadataCatalogConfigGlueDataCatalog struct {
	// A string containing the value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appflow_flow#database_name AppflowFlow#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// A string containing the value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appflow_flow#role_arn AppflowFlow#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A string containing the value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appflow_flow#table_prefix AppflowFlow#table_prefix}
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
}

