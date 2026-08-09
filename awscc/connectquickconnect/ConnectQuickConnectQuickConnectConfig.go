// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectquickconnect


type ConnectQuickConnectQuickConnectConfig struct {
	// The type of quick connect.
	//
	// In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_quick_connect#quick_connect_type ConnectQuickConnect#quick_connect_type}
	QuickConnectType *string `field:"required" json:"quickConnectType" yaml:"quickConnectType"`
	// The flow configuration. This is required only if QuickConnectType is FLOW.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_quick_connect#flow_config ConnectQuickConnect#flow_config}
	FlowConfig *ConnectQuickConnectQuickConnectConfigFlowConfig `field:"optional" json:"flowConfig" yaml:"flowConfig"`
	// The phone configuration. This is required only if QuickConnectType is PHONE_NUMBER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_quick_connect#phone_config ConnectQuickConnect#phone_config}
	PhoneConfig *ConnectQuickConnectQuickConnectConfigPhoneConfig `field:"optional" json:"phoneConfig" yaml:"phoneConfig"`
	// The queue configuration. This is required only if QuickConnectType is QUEUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_quick_connect#queue_config ConnectQuickConnect#queue_config}
	QueueConfig *ConnectQuickConnectQuickConnectConfigQueueConfig `field:"optional" json:"queueConfig" yaml:"queueConfig"`
	// The user configuration. This is required only if QuickConnectType is USER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_quick_connect#user_config ConnectQuickConnect#user_config}
	UserConfig *ConnectQuickConnectQuickConnectConfigUserConfig `field:"optional" json:"userConfig" yaml:"userConfig"`
}

