// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentEksConfigurationAccessEntry struct {
	// The desired state of the EKS access entry managed by AWS Batch. When omitted, AWS Batch applies INHERIT_FROM_CLUSTER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/batch_compute_environment#desired_state BatchComputeEnvironment#desired_state}
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
}

