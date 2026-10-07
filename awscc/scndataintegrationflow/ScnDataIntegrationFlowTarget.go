// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow


type ScnDataIntegrationFlowTarget struct {
	// The target type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#target_type ScnDataIntegrationFlow#target_type}
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// The dataset target configuration parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#dataset_target ScnDataIntegrationFlow#dataset_target}
	DatasetTarget *ScnDataIntegrationFlowTargetDatasetTarget `field:"optional" json:"datasetTarget" yaml:"datasetTarget"`
}

