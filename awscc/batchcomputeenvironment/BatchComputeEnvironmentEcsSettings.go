// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentEcsSettings struct {
	// The CloudWatch Container Insights setting applied to the Amazon ECS cluster that backs this compute environment.
	//
	// After you set this property, you can't revert it to the default (unset) state in which the setting is managed outside of AWS Batch. If you remove this property after previously setting it, AWS Batch treats the omission as DISABLED, because the underlying API has no way to unset the value. Because of this, if a stack rollback would return this property to its previous unset state, AWS Batch sets it to DISABLED instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/batch_compute_environment#container_insights BatchComputeEnvironment#container_insights}
	ContainerInsights *string `field:"optional" json:"containerInsights" yaml:"containerInsights"`
}

