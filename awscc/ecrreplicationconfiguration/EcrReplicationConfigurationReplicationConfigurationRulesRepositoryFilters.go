// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfigurationRulesRepositoryFilters struct {
	// The repository filter details.
	//
	// When the ``PREFIX_MATCH`` filter type is specified, this value is required and should be the repository name prefix to configure replication for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecr_replication_configuration#filter EcrReplicationConfiguration#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// The repository filter type.
	//
	// The only supported value is ``PREFIX_MATCH``, which is a repository name prefix specified with the ``filter`` parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecr_replication_configuration#filter_type EcrReplicationConfiguration#filter_type}
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
}

