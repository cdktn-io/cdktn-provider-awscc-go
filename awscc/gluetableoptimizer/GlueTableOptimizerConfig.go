// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetableoptimizer

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueTableOptimizerConfig struct {
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
	// The catalog ID of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#catalog_id GlueTableOptimizer#catalog_id}
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The name of the database. For Hive compatibility, this is folded to lowercase when it is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#database_name GlueTableOptimizer#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The table name. For Hive compatibility, this must be entirely lowercase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#table_name GlueTableOptimizer#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Specifies configuration details of a table optimizer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#table_optimizer_configuration GlueTableOptimizer#table_optimizer_configuration}
	TableOptimizerConfiguration *GlueTableOptimizerTableOptimizerConfiguration `field:"required" json:"tableOptimizerConfiguration" yaml:"tableOptimizerConfiguration"`
	// The type of table optimizer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#type GlueTableOptimizer#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

