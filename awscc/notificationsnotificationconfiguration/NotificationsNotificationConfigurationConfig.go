// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notificationsnotificationconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NotificationsNotificationConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/notifications_notification_configuration#description NotificationsNotificationConfiguration#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/notifications_notification_configuration#name NotificationsNotificationConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/notifications_notification_configuration#aggregation_duration NotificationsNotificationConfiguration#aggregation_duration}.
	AggregationDuration *string `field:"optional" json:"aggregationDuration" yaml:"aggregationDuration"`
	// A list of tags that are attached to the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/notifications_notification_configuration#tags NotificationsNotificationConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

