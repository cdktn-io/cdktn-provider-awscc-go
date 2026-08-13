// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedatasource


type DatazoneDataSourceConfigurationRedshiftRunConfiguration struct {
	// The data access role included in the configuration details of the Amazon Redshift data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_data_source#data_access_role DatazoneDataSource#data_access_role}
	DataAccessRole *string `field:"optional" json:"dataAccessRole" yaml:"dataAccessRole"`
	// The details of the credentials required to access an Amazon Redshift cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_data_source#redshift_credential_configuration DatazoneDataSource#redshift_credential_configuration}
	RedshiftCredentialConfiguration *DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftCredentialConfiguration `field:"optional" json:"redshiftCredentialConfiguration" yaml:"redshiftCredentialConfiguration"`
	// The details of the Amazon Redshift storage as part of the configuration of an Amazon Redshift data source run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_data_source#redshift_storage DatazoneDataSource#redshift_storage}
	RedshiftStorage *DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorage `field:"optional" json:"redshiftStorage" yaml:"redshiftStorage"`
	// The relational filter configurations included in the configuration details of the Amazon Redshift data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_data_source#relational_filter_configurations DatazoneDataSource#relational_filter_configurations}
	RelationalFilterConfigurations interface{} `field:"optional" json:"relationalFilterConfigurations" yaml:"relationalFilterConfigurations"`
}

