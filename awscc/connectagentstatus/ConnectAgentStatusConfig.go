// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectagentstatus

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectAgentStatusConfig struct {
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
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_agent_status#instance_arn ConnectAgentStatus#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_agent_status#name ConnectAgentStatus#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The state of the status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_agent_status#state ConnectAgentStatus#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// The description of the status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_agent_status#description ConnectAgentStatus#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display order of the status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_agent_status#display_order ConnectAgentStatus#display_order}
	DisplayOrder *float64 `field:"optional" json:"displayOrder" yaml:"displayOrder"`
	// A number indicating the reset order of the agent status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_agent_status#reset_order_number ConnectAgentStatus#reset_order_number}
	ResetOrderNumber interface{} `field:"optional" json:"resetOrderNumber" yaml:"resetOrderNumber"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_agent_status#tags ConnectAgentStatus#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of agent status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_agent_status#type ConnectAgentStatus#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

