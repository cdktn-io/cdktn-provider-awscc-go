// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow


type ScnDataIntegrationFlowTargetDatasetTargetOptions struct {
	// The option to perform deduplication on data records sharing same primary key values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/scn_data_integration_flow#dedupe_records ScnDataIntegrationFlow#dedupe_records}
	DedupeRecords interface{} `field:"optional" json:"dedupeRecords" yaml:"dedupeRecords"`
	// The deduplication strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/scn_data_integration_flow#dedupe_strategy ScnDataIntegrationFlow#dedupe_strategy}
	DedupeStrategy *ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategy `field:"optional" json:"dedupeStrategy" yaml:"dedupeStrategy"`
	// The load type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/scn_data_integration_flow#load_type ScnDataIntegrationFlow#load_type}
	LoadType *string `field:"optional" json:"loadType" yaml:"loadType"`
}

