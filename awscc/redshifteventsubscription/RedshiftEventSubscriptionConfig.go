// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshifteventsubscription

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftEventSubscriptionConfig struct {
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
	// The name of the Amazon Redshift event notification subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_event_subscription#subscription_name RedshiftEventSubscription#subscription_name}
	SubscriptionName *string `field:"required" json:"subscriptionName" yaml:"subscriptionName"`
	// A boolean value;
	//
	// set to true to activate the subscription, and set to false to create the subscription but not activate it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_event_subscription#enabled RedshiftEventSubscription#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the Amazon Redshift event categories to be published by the event notification subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_event_subscription#event_categories RedshiftEventSubscription#event_categories}
	EventCategories *[]*string `field:"optional" json:"eventCategories" yaml:"eventCategories"`
	// Specifies the Amazon Redshift event severity to be published by the event notification subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_event_subscription#severity RedshiftEventSubscription#severity}
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_event_subscription#sns_topic_arn RedshiftEventSubscription#sns_topic_arn}
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
	// A list of one or more identifiers of Amazon Redshift source objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_event_subscription#source_ids RedshiftEventSubscription#source_ids}
	SourceIds *[]*string `field:"optional" json:"sourceIds" yaml:"sourceIds"`
	// The type of source that will be generating the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_event_subscription#source_type RedshiftEventSubscription#source_type}
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_event_subscription#tags RedshiftEventSubscription#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

