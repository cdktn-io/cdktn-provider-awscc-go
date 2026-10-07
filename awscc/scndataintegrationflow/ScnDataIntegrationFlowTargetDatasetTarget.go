// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow


type ScnDataIntegrationFlowTargetDatasetTarget struct {
	// The dataset ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/scn_data_integration_flow#dataset_identifier ScnDataIntegrationFlow#dataset_identifier}
	DatasetIdentifier *string `field:"optional" json:"datasetIdentifier" yaml:"datasetIdentifier"`
	// The dataset options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/scn_data_integration_flow#options ScnDataIntegrationFlow#options}
	Options *ScnDataIntegrationFlowTargetDatasetTargetOptions `field:"optional" json:"options" yaml:"options"`
}

