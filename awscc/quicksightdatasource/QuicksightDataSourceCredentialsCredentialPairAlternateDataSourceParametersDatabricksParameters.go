// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersDatabricksParameters struct {
	// <p>The host name of the Databricks data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// <p>The port for the Databricks data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_source#port QuicksightDataSource#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// <p>The HTTP path of the Databricks data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_source#sql_endpoint_path QuicksightDataSource#sql_endpoint_path}
	SqlEndpointPath *string `field:"optional" json:"sqlEndpointPath" yaml:"sqlEndpointPath"`
}

