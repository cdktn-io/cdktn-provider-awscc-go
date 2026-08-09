// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmlconfiguredmodelalgorithmassociation


type CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModelInferenceJobsContainerLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanroomsml_configured_model_algorithm_association#allowed_account_ids CleanroomsmlConfiguredModelAlgorithmAssociation#allowed_account_ids}.
	AllowedAccountIds *[]*string `field:"optional" json:"allowedAccountIds" yaml:"allowedAccountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanroomsml_configured_model_algorithm_association#filter_pattern CleanroomsmlConfiguredModelAlgorithmAssociation#filter_pattern}.
	FilterPattern *string `field:"optional" json:"filterPattern" yaml:"filterPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanroomsml_configured_model_algorithm_association#log_redaction_configuration CleanroomsmlConfiguredModelAlgorithmAssociation#log_redaction_configuration}.
	LogRedactionConfiguration *CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfigurationPoliciesTrainedModelInferenceJobsContainerLogsLogRedactionConfiguration `field:"optional" json:"logRedactionConfiguration" yaml:"logRedactionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanroomsml_configured_model_algorithm_association#log_type CleanroomsmlConfiguredModelAlgorithmAssociation#log_type}.
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
}

