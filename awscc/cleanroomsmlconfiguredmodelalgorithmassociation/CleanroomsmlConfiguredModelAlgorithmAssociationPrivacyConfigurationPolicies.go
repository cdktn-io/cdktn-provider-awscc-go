// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmlconfiguredmodelalgorithmassociation


type CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanroomsml_configured_model_algorithm_association#trained_model_exports CleanroomsmlConfiguredModelAlgorithmAssociation#trained_model_exports}.
	TrainedModelExports *CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModelExports `field:"optional" json:"trainedModelExports" yaml:"trainedModelExports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanroomsml_configured_model_algorithm_association#trained_model_inference_jobs CleanroomsmlConfiguredModelAlgorithmAssociation#trained_model_inference_jobs}.
	TrainedModelInferenceJobs *CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModelInferenceJobs `field:"optional" json:"trainedModelInferenceJobs" yaml:"trainedModelInferenceJobs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanroomsml_configured_model_algorithm_association#trained_models CleanroomsmlConfiguredModelAlgorithmAssociation#trained_models}.
	TrainedModels *CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModels `field:"optional" json:"trainedModels" yaml:"trainedModels"`
}

