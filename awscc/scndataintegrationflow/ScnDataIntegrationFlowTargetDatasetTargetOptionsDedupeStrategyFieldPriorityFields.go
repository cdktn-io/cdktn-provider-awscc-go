// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow


type ScnDataIntegrationFlowTargetDatasetTargetOptionsDedupeStrategyFieldPriorityFields struct {
	// The name of the deduplication field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#name ScnDataIntegrationFlow#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sort order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#sort_order ScnDataIntegrationFlow#sort_order}
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
}

