// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchAlertConfig struct {
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
	// The display name of the alert.
	//
	// Names don't need to be unique within a space; the alert is identified by AlertId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#name CloudwatchAlert#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the access profile the alert uses to run its query and send notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#profile_id CloudwatchAlert#profile_id}
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// The rule that defines how the alert is evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#rule CloudwatchAlert#rule}
	Rule *CloudwatchAlertRule `field:"required" json:"rule" yaml:"rule"`
	// The ID of the space to create the alert in. Changing this value replaces the alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#space_id CloudwatchAlert#space_id}
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
	// A description of the alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#description CloudwatchAlert#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The notification rules that determine when and where notifications are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#notification_rules CloudwatchAlert#notification_rules}
	NotificationRules interface{} `field:"optional" json:"notificationRules" yaml:"notificationRules"`
	// Whether the alert sends notifications.
	//
	// When false, the alert is still evaluated and records its state, but sends nothing. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#notifications_enabled CloudwatchAlert#notifications_enabled}
	NotificationsEnabled interface{} `field:"optional" json:"notificationsEnabled" yaml:"notificationsEnabled"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#tags CloudwatchAlert#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

