// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedatasource


type DatazoneDataSourceConfigurationGlueRunConfiguration struct {
	// Specifies whether to automatically import data quality metrics as part of the data source run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_data_source#auto_import_data_quality_result DatazoneDataSource#auto_import_data_quality_result}
	AutoImportDataQualityResult interface{} `field:"optional" json:"autoImportDataQualityResult" yaml:"autoImportDataQualityResult"`
	// The catalog name in the AWS Glue run configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_data_source#catalog_name DatazoneDataSource#catalog_name}
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// The data access role included in the configuration details of the AWS Glue data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_data_source#data_access_role DatazoneDataSource#data_access_role}
	DataAccessRole *string `field:"optional" json:"dataAccessRole" yaml:"dataAccessRole"`
	// The relational filter configurations included in the configuration details of the AWS Glue data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_data_source#relational_filter_configurations DatazoneDataSource#relational_filter_configurations}
	RelationalFilterConfigurations interface{} `field:"optional" json:"relationalFilterConfigurations" yaml:"relationalFilterConfigurations"`
}

