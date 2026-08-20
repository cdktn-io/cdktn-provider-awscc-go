// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubapp


type ResiliencehubAppEventSubscriptions struct {
	// The type of event you would like to subscribe and get notification for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/resiliencehub_app#event_type ResiliencehubApp#event_type}
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// Unique name to identify an event subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/resiliencehub_app#name ResiliencehubApp#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/resiliencehub_app#sns_topic_arn ResiliencehubApp#sns_topic_arn}
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

