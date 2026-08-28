// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmlconfiguredmodelalgorithmassociation


type CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModelExports struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm_association#files_to_export CleanroomsmlConfiguredModelAlgorithmAssociation#files_to_export}.
	FilesToExport *[]*string `field:"optional" json:"filesToExport" yaml:"filesToExport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm_association#max_size CleanroomsmlConfiguredModelAlgorithmAssociation#max_size}.
	MaxSize *CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModelExportsMaxSize `field:"optional" json:"maxSize" yaml:"maxSize"`
}

