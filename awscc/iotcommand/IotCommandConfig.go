// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotcommand

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotCommandConfig struct {
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
	// The unique identifier for the command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#command_id IotCommand#command_id}
	CommandId *string `field:"required" json:"commandId" yaml:"commandId"`
	// The date and time when the command was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#created_at IotCommand#created_at}
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// A flag indicating whether the command is deprecated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#deprecated IotCommand#deprecated}
	Deprecated interface{} `field:"optional" json:"deprecated" yaml:"deprecated"`
	// The description of the command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#description IotCommand#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name for the command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#display_name IotCommand#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The date and time when the command was last updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#last_updated_at IotCommand#last_updated_at}
	LastUpdatedAt *string `field:"optional" json:"lastUpdatedAt" yaml:"lastUpdatedAt"`
	// The list of mandatory parameters for the command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#mandatory_parameters IotCommand#mandatory_parameters}
	MandatoryParameters interface{} `field:"optional" json:"mandatoryParameters" yaml:"mandatoryParameters"`
	// The namespace to which the command belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#namespace IotCommand#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The payload associated with the command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#payload IotCommand#payload}
	Payload *IotCommandPayload `field:"optional" json:"payload" yaml:"payload"`
	// The payload template associated with the command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#payload_template IotCommand#payload_template}
	PayloadTemplate *string `field:"optional" json:"payloadTemplate" yaml:"payloadTemplate"`
	// A flag indicating whether the command is pending deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#pending_deletion IotCommand#pending_deletion}
	PendingDeletion interface{} `field:"optional" json:"pendingDeletion" yaml:"pendingDeletion"`
	// The command preprocessor configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#preprocessor IotCommand#preprocessor}
	Preprocessor *IotCommandPreprocessor `field:"optional" json:"preprocessor" yaml:"preprocessor"`
	// The customer role associated with the command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#role_arn IotCommand#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags to be associated with the command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_command#tags IotCommand#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

