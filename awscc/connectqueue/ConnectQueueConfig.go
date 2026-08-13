// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectqueue

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectQueueConfig struct {
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
	// The identifier for the hours of operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_queue#hours_of_operation_arn ConnectQueue#hours_of_operation_arn}
	HoursOfOperationArn *string `field:"required" json:"hoursOfOperationArn" yaml:"hoursOfOperationArn"`
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_queue#instance_arn ConnectQueue#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_queue#name ConnectQueue#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The email addresses that agents can use when replying to or initiating email contacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_queue#additional_email_addresses ConnectQueue#additional_email_addresses}
	AdditionalEmailAddresses interface{} `field:"optional" json:"additionalEmailAddresses" yaml:"additionalEmailAddresses"`
	// The description of the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_queue#description ConnectQueue#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The maximum number of contacts that can be in the queue before it is considered full.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_queue#max_contacts ConnectQueue#max_contacts}
	MaxContacts *float64 `field:"optional" json:"maxContacts" yaml:"maxContacts"`
	// The outbound caller ID name, number, and outbound whisper flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_queue#outbound_caller_config ConnectQueue#outbound_caller_config}
	OutboundCallerConfig *ConnectQueueOutboundCallerConfig `field:"optional" json:"outboundCallerConfig" yaml:"outboundCallerConfig"`
	// The outbound email address ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_queue#outbound_email_config ConnectQueue#outbound_email_config}
	OutboundEmailConfig *ConnectQueueOutboundEmailConfig `field:"optional" json:"outboundEmailConfig" yaml:"outboundEmailConfig"`
	// The quick connects available to agents who are working the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_queue#quick_connect_arns ConnectQueue#quick_connect_arns}
	QuickConnectArns *[]*string `field:"optional" json:"quickConnectArns" yaml:"quickConnectArns"`
	// The status of the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_queue#status ConnectQueue#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_queue#tags ConnectQueue#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

