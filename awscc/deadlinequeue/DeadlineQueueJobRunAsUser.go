// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeue


type DeadlineQueueJobRunAsUser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/deadline_queue#posix DeadlineQueue#posix}.
	Posix *DeadlineQueueJobRunAsUserPosix `field:"optional" json:"posix" yaml:"posix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/deadline_queue#run_as DeadlineQueue#run_as}.
	RunAs *string `field:"optional" json:"runAs" yaml:"runAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/deadline_queue#windows DeadlineQueue#windows}.
	Windows *DeadlineQueueJobRunAsUserWindows `field:"optional" json:"windows" yaml:"windows"`
}

