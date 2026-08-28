// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsdestination

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogsDestinationConfig struct {
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
	// The name of the destination resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_destination#destination_name LogsDestination#destination_name}
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
	// The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_destination#role_arn LogsDestination#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The ARN of the physical target where the log events are delivered (for example, a Kinesis stream).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_destination#target_arn LogsDestination#target_arn}
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_destination#destination_policy LogsDestination#destination_policy}
	DestinationPolicy *string `field:"optional" json:"destinationPolicy" yaml:"destinationPolicy"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_destination#tags LogsDestination#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

