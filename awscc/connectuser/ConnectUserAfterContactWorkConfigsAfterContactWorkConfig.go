// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuser


type ConnectUserAfterContactWorkConfigsAfterContactWorkConfig struct {
	// The after contact work (ACW) mode for the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_user#after_contact_work_mode ConnectUser#after_contact_work_mode}
	AfterContactWorkMode *string `field:"optional" json:"afterContactWorkMode" yaml:"afterContactWorkMode"`
	// The After Call Work (ACW) timeout setting, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_user#after_contact_work_time_limit ConnectUser#after_contact_work_time_limit}
	AfterContactWorkTimeLimit *float64 `field:"optional" json:"afterContactWorkTimeLimit" yaml:"afterContactWorkTimeLimit"`
}

