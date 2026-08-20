// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchdatasource


type OpensearchDataSourceDataSourceType struct {
	// Configuration for an S3 Glue Data Catalog data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearch_data_source#s3_glue_data_catalog OpensearchDataSource#s3_glue_data_catalog}
	S3GlueDataCatalog *OpensearchDataSourceDataSourceTypeS3GlueDataCatalog `field:"optional" json:"s3GlueDataCatalog" yaml:"s3GlueDataCatalog"`
}

