// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrrepositorycreationtemplate


type EcrRepositoryCreationTemplateImageTagMutabilityExclusionFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecr_repository_creation_template#image_tag_mutability_exclusion_filter_type EcrRepositoryCreationTemplate#image_tag_mutability_exclusion_filter_type}.
	ImageTagMutabilityExclusionFilterType *string `field:"optional" json:"imageTagMutabilityExclusionFilterType" yaml:"imageTagMutabilityExclusionFilterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecr_repository_creation_template#image_tag_mutability_exclusion_filter_value EcrRepositoryCreationTemplate#image_tag_mutability_exclusion_filter_value}.
	ImageTagMutabilityExclusionFilterValue *string `field:"optional" json:"imageTagMutabilityExclusionFilterValue" yaml:"imageTagMutabilityExclusionFilterValue"`
}

