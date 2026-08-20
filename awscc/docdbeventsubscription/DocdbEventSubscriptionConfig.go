// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package docdbeventsubscription

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DocdbEventSubscriptionConfig struct {
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
	// The Amazon Resource Name (ARN) of the SNS topic created for event notification.
	//
	// Amazon SNS creates the ARN when you create a topic and subscribe to it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/docdb_event_subscription#sns_topic_arn DocdbEventSubscription#sns_topic_arn}
	SnsTopicArn *string `field:"required" json:"snsTopicArn" yaml:"snsTopicArn"`
	// A Boolean value;
	//
	// set to true to activate the subscription, set to false to create the subscription but not active it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/docdb_event_subscription#enabled DocdbEventSubscription#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of event categories for a SourceType that you want to subscribe to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/docdb_event_subscription#event_categories DocdbEventSubscription#event_categories}
	EventCategories *[]*string `field:"optional" json:"eventCategories" yaml:"eventCategories"`
	// The list of identifiers of the event sources for which events are returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/docdb_event_subscription#source_ids DocdbEventSubscription#source_ids}
	SourceIds *[]*string `field:"optional" json:"sourceIds" yaml:"sourceIds"`
	// The type of source that is generating the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/docdb_event_subscription#source_type DocdbEventSubscription#source_type}
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// The name of the subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/docdb_event_subscription#subscription_name DocdbEventSubscription#subscription_name}
	SubscriptionName *string `field:"optional" json:"subscriptionName" yaml:"subscriptionName"`
}

