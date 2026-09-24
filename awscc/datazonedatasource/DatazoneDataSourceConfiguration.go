// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedatasource


type DatazoneDataSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_data_source#glue_run_configuration DatazoneDataSource#glue_run_configuration}.
	GlueRunConfiguration *DatazoneDataSourceConfigurationGlueRunConfiguration `field:"optional" json:"glueRunConfiguration" yaml:"glueRunConfiguration"`
	// The configuration details of the Amazon Redshift data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_data_source#redshift_run_configuration DatazoneDataSource#redshift_run_configuration}
	RedshiftRunConfiguration *DatazoneDataSourceConfigurationRedshiftRunConfiguration `field:"optional" json:"redshiftRunConfiguration" yaml:"redshiftRunConfiguration"`
	// The configuration details of the Amazon SageMaker data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_data_source#sage_maker_run_configuration DatazoneDataSource#sage_maker_run_configuration}
	SageMakerRunConfiguration *DatazoneDataSourceConfigurationSageMakerRunConfiguration `field:"optional" json:"sageMakerRunConfiguration" yaml:"sageMakerRunConfiguration"`
}

