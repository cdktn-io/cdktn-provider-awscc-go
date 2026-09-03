// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchdatasource


type OpensearchDataSourceDataSourceTypeS3GlueDataCatalog struct {
	// The ARN of the IAM role that grants OpenSearch Service permission to access the Glue Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearch_data_source#role_arn OpensearchDataSource#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

