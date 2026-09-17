// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagenttrigger


type DevopsagentTriggerCondition struct {
	// Schedule configuration for a time-based trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/devopsagent_trigger#schedule DevopsagentTrigger#schedule}
	Schedule *DevopsagentTriggerConditionSchedule `field:"required" json:"schedule" yaml:"schedule"`
}

