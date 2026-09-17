// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsanalysistemplate


type CleanroomsAnalysisTemplateSourceMetadataArtifacts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanrooms_analysis_template#additional_artifact_hashes CleanroomsAnalysisTemplate#additional_artifact_hashes}.
	AdditionalArtifactHashes interface{} `field:"optional" json:"additionalArtifactHashes" yaml:"additionalArtifactHashes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanrooms_analysis_template#entry_point_hash CleanroomsAnalysisTemplate#entry_point_hash}.
	EntryPointHash *CleanroomsAnalysisTemplateSourceMetadataArtifactsEntryPointHash `field:"optional" json:"entryPointHash" yaml:"entryPointHash"`
}

