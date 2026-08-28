// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetrigger


type GlueTriggerActionsNotificationProperty struct {
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_trigger#notify_delay_after GlueTrigger#notify_delay_after}
	NotifyDelayAfter *float64 `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
}

