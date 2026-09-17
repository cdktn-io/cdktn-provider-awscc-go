// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagenttrigger


type DevopsagentTriggerConditionSchedule struct {
	// A cron or rate expression that defines when the trigger fires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/devopsagent_trigger#expression DevopsagentTrigger#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

