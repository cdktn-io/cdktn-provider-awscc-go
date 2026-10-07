// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow


type ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategy struct {
	// The field priority deduplication strategy configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#field_priority ScnDataIntegrationFlow#field_priority}
	FieldPriority *ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriority `field:"optional" json:"fieldPriority" yaml:"fieldPriority"`
	// The deduplication strategy type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#type ScnDataIntegrationFlow#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

