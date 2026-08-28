// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup


type SagemakerFeatureGroupOfflineStoreConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_feature_group#data_catalog_config SagemakerFeatureGroup#data_catalog_config}.
	DataCatalogConfig *SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig `field:"optional" json:"dataCatalogConfig" yaml:"dataCatalogConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_feature_group#disable_glue_table_creation SagemakerFeatureGroup#disable_glue_table_creation}.
	DisableGlueTableCreation interface{} `field:"optional" json:"disableGlueTableCreation" yaml:"disableGlueTableCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_feature_group#s3_storage_config SagemakerFeatureGroup#s3_storage_config}.
	S3StorageConfig *SagemakerFeatureGroupOfflineStoreConfigS3StorageConfig `field:"optional" json:"s3StorageConfig" yaml:"s3StorageConfig"`
	// Format for the offline store feature group.
	//
	// Iceberg is the optimal format for feature groups shared between offline and online stores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_feature_group#table_format SagemakerFeatureGroup#table_format}
	TableFormat *string `field:"optional" json:"tableFormat" yaml:"tableFormat"`
}

