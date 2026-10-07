// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow


type ScnDataIntegrationFlowSources struct {
	// The source name that can be used as table alias in SQL transformation query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/scn_data_integration_flow#source_name ScnDataIntegrationFlow#source_name}
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// The source type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/scn_data_integration_flow#source_type ScnDataIntegrationFlow#source_type}
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// The dataset source configuration parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/scn_data_integration_flow#dataset_source ScnDataIntegrationFlow#dataset_source}
	DatasetSource *ScnDataIntegrationFlowSourcesDatasetSource `field:"optional" json:"datasetSource" yaml:"datasetSource"`
	// The S3 source configuration parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/scn_data_integration_flow#s3_source ScnDataIntegrationFlow#s3_source}
	S3Source *ScnDataIntegrationFlowSourcesS3Source `field:"optional" json:"s3Source" yaml:"s3Source"`
}

