// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentEcsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/batch_compute_environment#container_insights BatchComputeEnvironment#container_insights}.
	ContainerInsights *string `field:"optional" json:"containerInsights" yaml:"containerInsights"`
}

