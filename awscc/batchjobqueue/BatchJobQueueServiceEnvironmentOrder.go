// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobqueue


type BatchJobQueueServiceEnvironmentOrder struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/batch_job_queue#order BatchJobQueue#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/batch_job_queue#service_environment BatchJobQueue#service_environment}.
	ServiceEnvironment *string `field:"optional" json:"serviceEnvironment" yaml:"serviceEnvironment"`
}

