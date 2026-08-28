// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationdatacellsfilter

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LakeformationDataCellsFilterConfig struct {
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
	// The name of the Database that the Table resides in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_data_cells_filter#database_name LakeformationDataCellsFilter#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The desired name of the Data Cells Filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_data_cells_filter#name LakeformationDataCellsFilter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Catalog Id of the Table on which to create a Data Cells Filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_data_cells_filter#table_catalog_id LakeformationDataCellsFilter#table_catalog_id}
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// The name of the Table to create a Data Cells Filter for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_data_cells_filter#table_name LakeformationDataCellsFilter#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// A list of columns to be included in this Data Cells Filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_data_cells_filter#column_names LakeformationDataCellsFilter#column_names}
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_data_cells_filter#column_wildcard LakeformationDataCellsFilter#column_wildcard}
	ColumnWildcard *LakeformationDataCellsFilterColumnWildcard `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
	// An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_data_cells_filter#row_filter LakeformationDataCellsFilter#row_filter}
	RowFilter *LakeformationDataCellsFilterRowFilter `field:"optional" json:"rowFilter" yaml:"rowFilter"`
}

