// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsanalysistemplate


type CleanroomsAnalysisTemplateSourceArtifacts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_analysis_template#additional_artifacts CleanroomsAnalysisTemplate#additional_artifacts}.
	AdditionalArtifacts interface{} `field:"optional" json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_analysis_template#entry_point CleanroomsAnalysisTemplate#entry_point}.
	EntryPoint *CleanroomsAnalysisTemplateSourceArtifactsEntryPoint `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_analysis_template#role_arn CleanroomsAnalysisTemplate#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

