// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectquickconnect


type ConnectQuickConnectQuickConnectConfigQueueConfig struct {
	// The identifier of the contact flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_quick_connect#contact_flow_arn ConnectQuickConnect#contact_flow_arn}
	ContactFlowArn *string `field:"optional" json:"contactFlowArn" yaml:"contactFlowArn"`
	// The identifier for the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_quick_connect#queue_arn ConnectQuickConnect#queue_arn}
	QueueArn *string `field:"optional" json:"queueArn" yaml:"queueArn"`
}

