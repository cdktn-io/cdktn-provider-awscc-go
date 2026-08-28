// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitylakesubscribernotification

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecuritylakeSubscriberNotificationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securitylake_subscriber_notification#notification_configuration SecuritylakeSubscriberNotification#notification_configuration}.
	NotificationConfiguration *SecuritylakeSubscriberNotificationNotificationConfiguration `field:"required" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// The ARN for the subscriber.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securitylake_subscriber_notification#subscriber_arn SecuritylakeSubscriberNotification#subscriber_arn}
	SubscriberArn *string `field:"required" json:"subscriberArn" yaml:"subscriberArn"`
}

