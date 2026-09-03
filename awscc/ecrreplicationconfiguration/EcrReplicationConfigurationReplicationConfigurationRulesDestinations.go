// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfigurationRulesDestinations struct {
	// The Region to replicate to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecr_replication_configuration#region EcrReplicationConfiguration#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The AWS account ID of the Amazon ECR private registry to replicate to.
	//
	// When configuring cross-Region replication within your own registry, specify your own account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecr_replication_configuration#registry_id EcrReplicationConfiguration#registry_id}
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

