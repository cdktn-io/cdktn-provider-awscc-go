// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobqueue


type BatchJobQueueJobStateTimeLimitActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/batch_job_queue#action BatchJobQueue#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/batch_job_queue#max_time_seconds BatchJobQueue#max_time_seconds}.
	MaxTimeSeconds *float64 `field:"optional" json:"maxTimeSeconds" yaml:"maxTimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/batch_job_queue#reason BatchJobQueue#reason}.
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/batch_job_queue#state BatchJobQueue#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

