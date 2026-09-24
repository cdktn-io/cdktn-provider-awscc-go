// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetPhysicalTableMap struct {
	// <p>A physical table type built from the results of the custom SQL query.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_set#custom_sql QuicksightDataSet#custom_sql}
	CustomSql *QuicksightDataSetPhysicalTableMapCustomSql `field:"optional" json:"customSql" yaml:"customSql"`
	// <p>A physical table type for relational data sources.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_set#relational_table QuicksightDataSet#relational_table}
	RelationalTable *QuicksightDataSetPhysicalTableMapRelationalTable `field:"optional" json:"relationalTable" yaml:"relationalTable"`
	// <p>A physical table type for as S3 data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_data_set#s3_source QuicksightDataSet#s3_source}
	S3Source *QuicksightDataSetPhysicalTableMapS3Source `field:"optional" json:"s3Source" yaml:"s3Source"`
}

