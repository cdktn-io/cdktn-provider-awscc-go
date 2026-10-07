// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentagentspace


type DevopsagentAgentSpacePreferences struct {
	// Indicates whether elevated directed actions are permitted in this AgentSpace. Defaults to false when not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/devopsagent_agent_space#elevated_actions_enabled DevopsagentAgentSpace#elevated_actions_enabled}
	ElevatedActionsEnabled interface{} `field:"optional" json:"elevatedActionsEnabled" yaml:"elevatedActionsEnabled"`
}

