// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentComputeResourcesManagedInstancesProviderInstanceLaunchTemplateNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_compute_environment#security_groups BatchComputeEnvironment#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_compute_environment#subnets BatchComputeEnvironment#subnets}.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

