// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationcentralizationrule


type ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationTagPropagationConfiguration struct {
	// The ARN of the destination account IAM role used for tag propagation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/observabilityadmin_organization_centralization_rule#destination_role_arn ObservabilityadminOrganizationCentralizationRule#destination_role_arn}
	DestinationRoleArn *string `field:"optional" json:"destinationRoleArn" yaml:"destinationRoleArn"`
	// The strategy to resolve tag conflicts during propagation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/observabilityadmin_organization_centralization_rule#tag_conflict_resolution_strategy ObservabilityadminOrganizationCentralizationRule#tag_conflict_resolution_strategy}
	TagConflictResolutionStrategy *string `field:"optional" json:"tagConflictResolutionStrategy" yaml:"tagConflictResolutionStrategy"`
}

