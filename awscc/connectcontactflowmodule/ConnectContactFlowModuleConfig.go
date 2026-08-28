// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcontactflowmodule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectContactFlowModuleConfig struct {
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
	// The content of the contact flow module in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_contact_flow_module#content ConnectContactFlowModule#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The identifier of the Amazon Connect instance (ARN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_contact_flow_module#instance_arn ConnectContactFlowModule#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the contact flow module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_contact_flow_module#name ConnectContactFlowModule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the contact flow module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_contact_flow_module#description ConnectContactFlowModule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Defines the external invocation configuration of the flow module resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_contact_flow_module#external_invocation_configuration ConnectContactFlowModule#external_invocation_configuration}
	ExternalInvocationConfiguration *ConnectContactFlowModuleExternalInvocationConfiguration `field:"optional" json:"externalInvocationConfiguration" yaml:"externalInvocationConfiguration"`
	// The schema of the settings for contact flow module in JSON Schema V4 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_contact_flow_module#settings ConnectContactFlowModule#settings}
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
	// The state of the contact flow module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_contact_flow_module#state ConnectContactFlowModule#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_contact_flow_module#tags ConnectContactFlowModule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

