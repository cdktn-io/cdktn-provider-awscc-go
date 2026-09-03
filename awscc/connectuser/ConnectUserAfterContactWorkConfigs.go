// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuser


type ConnectUserAfterContactWorkConfigs struct {
	// After Contact Work configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user#after_contact_work_config ConnectUser#after_contact_work_config}
	AfterContactWorkConfig *ConnectUserAfterContactWorkConfigsAfterContactWorkConfig `field:"optional" json:"afterContactWorkConfig" yaml:"afterContactWorkConfig"`
	// After Contact Work configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user#agent_first_callback_after_contact_work_config ConnectUser#agent_first_callback_after_contact_work_config}
	AgentFirstCallbackAfterContactWorkConfig *ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfig `field:"optional" json:"agentFirstCallbackAfterContactWorkConfig" yaml:"agentFirstCallbackAfterContactWorkConfig"`
	// The channels that agents can handle in the Contact Control Panel (CCP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user#channel ConnectUser#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
}

