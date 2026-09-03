// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsanalysistemplate


type CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_analysis_template#column_classification CleanroomsAnalysisTemplate#column_classification}.
	ColumnClassification *CleanroomsAnalysisTemplateSyntheticDataParametersMlSyntheticDataParametersColumnClassification `field:"optional" json:"columnClassification" yaml:"columnClassification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_analysis_template#epsilon CleanroomsAnalysisTemplate#epsilon}.
	Epsilon *float64 `field:"optional" json:"epsilon" yaml:"epsilon"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_analysis_template#max_membership_inference_attack_score CleanroomsAnalysisTemplate#max_membership_inference_attack_score}.
	MaxMembershipInferenceAttackScore *float64 `field:"optional" json:"maxMembershipInferenceAttackScore" yaml:"maxMembershipInferenceAttackScore"`
}

