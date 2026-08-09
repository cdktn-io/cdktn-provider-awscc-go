// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob


type DatabrewJobDatabaseOutputs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#database_options DatabrewJob#database_options}.
	DatabaseOptions *DatabrewJobDatabaseOutputsDatabaseOptions `field:"optional" json:"databaseOptions" yaml:"databaseOptions"`
	// Database table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#database_output_mode DatabrewJob#database_output_mode}
	DatabaseOutputMode *string `field:"optional" json:"databaseOutputMode" yaml:"databaseOutputMode"`
	// Glue connection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#glue_connection_name DatabrewJob#glue_connection_name}
	GlueConnectionName *string `field:"optional" json:"glueConnectionName" yaml:"glueConnectionName"`
}

