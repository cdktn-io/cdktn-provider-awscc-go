// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseportal

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotsitewisePortalConfig struct {
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
	// The AWS administrator's contact email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_portal#portal_contact_email IotsitewisePortal#portal_contact_email}
	PortalContactEmail *string `field:"required" json:"portalContactEmail" yaml:"portalContactEmail"`
	// A friendly name for the portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_portal#portal_name IotsitewisePortal#portal_name}
	PortalName *string `field:"required" json:"portalName" yaml:"portalName"`
	// The ARN of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_portal#role_arn IotsitewisePortal#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal.
	//
	// You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_portal#alarms IotsitewisePortal#alarms}
	Alarms *IotsitewisePortalAlarms `field:"optional" json:"alarms" yaml:"alarms"`
	// The email address that sends alarm notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_portal#notification_sender_email IotsitewisePortal#notification_sender_email}
	NotificationSenderEmail *string `field:"optional" json:"notificationSenderEmail" yaml:"notificationSenderEmail"`
	// The service to use to authenticate users to the portal.
	//
	// Choose from SSO or IAM. You can't change this value after you create a portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_portal#portal_auth_mode IotsitewisePortal#portal_auth_mode}
	PortalAuthMode *string `field:"optional" json:"portalAuthMode" yaml:"portalAuthMode"`
	// A description for the portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_portal#portal_description IotsitewisePortal#portal_description}
	PortalDescription *string `field:"optional" json:"portalDescription" yaml:"portalDescription"`
	// The type of portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_portal#portal_type IotsitewisePortal#portal_type}
	PortalType *string `field:"optional" json:"portalType" yaml:"portalType"`
	// Map to associate detail of configuration related with a PortalType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_portal#portal_type_configuration IotsitewisePortal#portal_type_configuration}
	PortalTypeConfiguration interface{} `field:"optional" json:"portalTypeConfiguration" yaml:"portalTypeConfiguration"`
	// A list of key-value pairs that contain metadata for the portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_portal#tags IotsitewisePortal#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

