// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrsigningconfiguration


type EcrSigningConfigurationRulesRepositoryFilters struct {
	// Repository name pattern (supports '*' wildcard).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecr_signing_configuration#filter EcrSigningConfiguration#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Type of repository filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecr_signing_configuration#filter_type EcrSigningConfiguration#filter_type}
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
}

