// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuser


type ConnectUserAutoAcceptConfigs struct {
	// The agent first callback auto accept setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#agent_first_callback_auto_accept ConnectUser#agent_first_callback_auto_accept}
	AgentFirstCallbackAutoAccept interface{} `field:"optional" json:"agentFirstCallbackAutoAccept" yaml:"agentFirstCallbackAutoAccept"`
	// The Auto accept setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#auto_accept ConnectUser#auto_accept}
	AutoAccept interface{} `field:"optional" json:"autoAccept" yaml:"autoAccept"`
	// The channels that agents can handle in the Contact Control Panel (CCP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#channel ConnectUser#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
}

