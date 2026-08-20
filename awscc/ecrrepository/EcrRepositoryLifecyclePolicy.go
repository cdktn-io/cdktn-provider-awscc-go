// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrrepository


type EcrRepositoryLifecyclePolicy struct {
	// The JSON repository policy text to apply to the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecr_repository#lifecycle_policy_text EcrRepository#lifecycle_policy_text}
	LifecyclePolicyText *string `field:"optional" json:"lifecyclePolicyText" yaml:"lifecyclePolicyText"`
	// The AWS account ID associated with the registry that contains the repository.
	//
	// If you do? not specify a registry, the default registry is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecr_repository#registry_id EcrRepository#registry_id}
	RegistryId *string `field:"optional" json:"registryId" yaml:"registryId"`
}

