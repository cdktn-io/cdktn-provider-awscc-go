// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmlconfiguredmodelalgorithmassociation


type CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanroomsml_configured_model_algorithm_association#container_logs CleanroomsmlConfiguredModelAlgorithmAssociation#container_logs}.
	ContainerLogs interface{} `field:"optional" json:"containerLogs" yaml:"containerLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanroomsml_configured_model_algorithm_association#container_metrics CleanroomsmlConfiguredModelAlgorithmAssociation#container_metrics}.
	ContainerMetrics *CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModelsContainerMetrics `field:"optional" json:"containerMetrics" yaml:"containerMetrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanroomsml_configured_model_algorithm_association#max_artifact_size CleanroomsmlConfiguredModelAlgorithmAssociation#max_artifact_size}.
	MaxArtifactSize *CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModelsMaxArtifactSize `field:"optional" json:"maxArtifactSize" yaml:"maxArtifactSize"`
}

