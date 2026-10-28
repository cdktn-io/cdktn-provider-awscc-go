// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueTableConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the Data Catalog in which to create the Table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#catalog_id GlueTable#catalog_id}
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The name of the database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#database_name GlueTable#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The table name. For Hive compatibility, this is folded to lowercase when it is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#name GlueTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies an OpenTableFormatInput structure when creating an open format table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#open_table_format_input GlueTable#open_table_format_input}
	OpenTableFormatInput *GlueTableOpenTableFormatInput `field:"optional" json:"openTableFormatInput" yaml:"openTableFormatInput"`
	// A structure used to define a table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#table_input GlueTable#table_input}
	TableInput *GlueTableTableInput `field:"optional" json:"tableInput" yaml:"tableInput"`
}

