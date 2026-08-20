// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrrepository


type EcrRepositoryImageTagMutabilityExclusionFilters struct {
	// Specifies the type of filter to use for excluding image tags from the repository's mutability setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecr_repository#image_tag_mutability_exclusion_filter_type EcrRepository#image_tag_mutability_exclusion_filter_type}
	ImageTagMutabilityExclusionFilterType *string `field:"optional" json:"imageTagMutabilityExclusionFilterType" yaml:"imageTagMutabilityExclusionFilterType"`
	// The value to use when filtering image tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecr_repository#image_tag_mutability_exclusion_filter_value EcrRepository#image_tag_mutability_exclusion_filter_value}
	ImageTagMutabilityExclusionFilterValue *string `field:"optional" json:"imageTagMutabilityExclusionFilterValue" yaml:"imageTagMutabilityExclusionFilterValue"`
}

