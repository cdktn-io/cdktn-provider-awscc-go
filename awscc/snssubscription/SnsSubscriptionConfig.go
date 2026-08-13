// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package snssubscription

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SnsSubscriptionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The subscription's protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sns_subscription#protocol SnsSubscription#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The ARN of the topic to subscribe to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sns_subscription#topic_arn SnsSubscription#topic_arn}
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// The delivery policy JSON assigned to the subscription.
	//
	// Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sns_subscription#delivery_policy SnsSubscription#delivery_policy}
	DeliveryPolicy *string `field:"optional" json:"deliveryPolicy" yaml:"deliveryPolicy"`
	// The subscription's endpoint. The endpoint value depends on the protocol that you specify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sns_subscription#endpoint SnsSubscription#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The filter policy JSON assigned to the subscription. Enables the subscriber to filter out unwanted messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sns_subscription#filter_policy SnsSubscription#filter_policy}
	FilterPolicy *string `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// This attribute lets you choose the filtering scope by using one of the following string value types: MessageAttributes (default) and MessageBody.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sns_subscription#filter_policy_scope SnsSubscription#filter_policy_scope}
	FilterPolicyScope *string `field:"optional" json:"filterPolicyScope" yaml:"filterPolicyScope"`
	// When set to true, enables raw message delivery.
	//
	// Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sns_subscription#raw_message_delivery SnsSubscription#raw_message_delivery}
	RawMessageDelivery interface{} `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue.
	//
	// Messages that can't be delivered due to client errors are held in the dead-letter queue for further analysis or reprocessing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sns_subscription#redrive_policy SnsSubscription#redrive_policy}
	RedrivePolicy *string `field:"optional" json:"redrivePolicy" yaml:"redrivePolicy"`
	// For cross-region subscriptions, the region in which the topic resides.If no region is specified, AWS CloudFormation uses the region of the caller as the default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sns_subscription#region SnsSubscription#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Specifies whether Amazon SNS resends the notification to the subscription when a message's attribute changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sns_subscription#replay_policy SnsSubscription#replay_policy}
	ReplayPolicy *string `field:"optional" json:"replayPolicy" yaml:"replayPolicy"`
	// This property applies only to Amazon Data Firehose delivery stream subscriptions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sns_subscription#subscription_role_arn SnsSubscription#subscription_role_arn}
	SubscriptionRoleArn *string `field:"optional" json:"subscriptionRoleArn" yaml:"subscriptionRoleArn"`
}

