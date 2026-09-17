// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobqueue


type BatchJobQueueComputeEnvironmentOrder struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/batch_job_queue#compute_environment BatchJobQueue#compute_environment}.
	ComputeEnvironment *string `field:"optional" json:"computeEnvironment" yaml:"computeEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/batch_job_queue#order BatchJobQueue#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
}

