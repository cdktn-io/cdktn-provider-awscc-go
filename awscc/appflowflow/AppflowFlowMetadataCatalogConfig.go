// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowMetadataCatalogConfig struct {
	// Configurations of glue data catalog of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appflow_flow#glue_data_catalog AppflowFlow#glue_data_catalog}
	GlueDataCatalog *AppflowFlowMetadataCatalogConfigGlueDataCatalog `field:"optional" json:"glueDataCatalog" yaml:"glueDataCatalog"`
}

