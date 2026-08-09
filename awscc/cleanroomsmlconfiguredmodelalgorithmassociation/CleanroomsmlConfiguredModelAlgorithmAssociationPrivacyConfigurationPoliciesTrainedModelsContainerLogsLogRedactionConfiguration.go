// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmlconfiguredmodelalgorithmassociation


type CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModelsContainerLogsLogRedactionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanroomsml_configured_model_algorithm_association#custom_entity_config CleanroomsmlConfiguredModelAlgorithmAssociation#custom_entity_config}.
	CustomEntityConfig *CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModelsContainerLogsLogRedactionConfigurationCustomEntityConfig `field:"optional" json:"customEntityConfig" yaml:"customEntityConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanroomsml_configured_model_algorithm_association#entities_to_redact CleanroomsmlConfiguredModelAlgorithmAssociation#entities_to_redact}.
	EntitiesToRedact *[]*string `field:"optional" json:"entitiesToRedact" yaml:"entitiesToRedact"`
}

