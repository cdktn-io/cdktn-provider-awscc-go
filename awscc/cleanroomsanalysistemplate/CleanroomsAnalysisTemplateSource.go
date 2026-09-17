// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsanalysistemplate


type CleanroomsAnalysisTemplateSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanrooms_analysis_template#artifacts CleanroomsAnalysisTemplate#artifacts}.
	Artifacts *CleanroomsAnalysisTemplateSourceArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanrooms_analysis_template#text CleanroomsAnalysisTemplate#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

