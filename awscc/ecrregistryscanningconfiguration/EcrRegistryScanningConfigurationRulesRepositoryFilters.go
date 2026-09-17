// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrregistryscanningconfiguration


type EcrRegistryScanningConfigurationRulesRepositoryFilters struct {
	// The filter to use when scanning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecr_registry_scanning_configuration#filter EcrRegistryScanningConfiguration#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// The type associated with the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecr_registry_scanning_configuration#filter_type EcrRegistryScanningConfiguration#filter_type}
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

