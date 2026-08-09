// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketNotificationConfiguration struct {
	// Enables delivery of events to Amazon EventBridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_bucket#event_bridge_configuration S3Bucket#event_bridge_configuration}
	EventBridgeConfiguration *S3BucketNotificationConfigurationEventBridgeConfiguration `field:"optional" json:"eventBridgeConfiguration" yaml:"eventBridgeConfiguration"`
	// Describes the LAMlong functions to invoke and the events for which to invoke them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_bucket#lambda_configurations S3Bucket#lambda_configurations}
	LambdaConfigurations interface{} `field:"optional" json:"lambdaConfigurations" yaml:"lambdaConfigurations"`
	// The Amazon Simple Queue Service queues to publish messages to and the events for which to publish messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_bucket#queue_configurations S3Bucket#queue_configurations}
	QueueConfigurations interface{} `field:"optional" json:"queueConfigurations" yaml:"queueConfigurations"`
	// The topic to which notifications are sent and the events for which notifications are generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_bucket#topic_configurations S3Bucket#topic_configurations}
	TopicConfigurations interface{} `field:"optional" json:"topicConfigurations" yaml:"topicConfigurations"`
}

