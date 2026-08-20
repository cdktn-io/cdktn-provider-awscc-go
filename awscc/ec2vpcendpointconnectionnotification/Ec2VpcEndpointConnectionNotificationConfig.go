// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpcendpointconnectionnotification

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VpcEndpointConnectionNotificationConfig struct {
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
	// The endpoint events for which to receive notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc_endpoint_connection_notification#connection_events Ec2VpcEndpointConnectionNotification#connection_events}
	ConnectionEvents *[]*string `field:"required" json:"connectionEvents" yaml:"connectionEvents"`
	// The ARN of the SNS topic for the notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc_endpoint_connection_notification#connection_notification_arn Ec2VpcEndpointConnectionNotification#connection_notification_arn}
	ConnectionNotificationArn *string `field:"required" json:"connectionNotificationArn" yaml:"connectionNotificationArn"`
	// The ID of the endpoint service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc_endpoint_connection_notification#service_id Ec2VpcEndpointConnectionNotification#service_id}
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
	// The ID of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_vpc_endpoint_connection_notification#vpc_endpoint_id Ec2VpcEndpointConnectionNotification#vpc_endpoint_id}
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

