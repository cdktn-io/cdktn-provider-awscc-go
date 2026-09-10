// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeue


type DeadlineQueueJobRunAsUserWindows struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/deadline_queue#password_arn DeadlineQueue#password_arn}.
	PasswordArn *string `field:"optional" json:"passwordArn" yaml:"passwordArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/deadline_queue#user DeadlineQueue#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

