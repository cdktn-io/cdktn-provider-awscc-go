// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetPhysicalTableMapCustomSql struct {
	// <p>The column schema from the SQL query result set.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_set#columns QuicksightDataSet#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// <p>The Amazon Resource Name (ARN) of the data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_set#data_source_arn QuicksightDataSet#data_source_arn}
	DataSourceArn *string `field:"optional" json:"dataSourceArn" yaml:"dataSourceArn"`
	// <p>A display name for the SQL query result.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_set#name QuicksightDataSet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// <p>The SQL query.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_set#sql_query QuicksightDataSet#sql_query}
	SqlQuery *string `field:"optional" json:"sqlQuery" yaml:"sqlQuery"`
}

