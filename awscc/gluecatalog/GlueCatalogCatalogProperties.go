// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog


type GlueCatalogCatalogProperties struct {
	// Data lake access properties for the catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_catalog#data_lake_access_properties GlueCatalog#data_lake_access_properties}
	DataLakeAccessProperties *GlueCatalogCatalogPropertiesDataLakeAccessProperties `field:"optional" json:"dataLakeAccessProperties" yaml:"dataLakeAccessProperties"`
}

