// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecttasktemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectTaskTemplateConfig struct {
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
	// The identifier (arn) of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_task_template#instance_arn ConnectTaskTemplate#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// the client token string in uuid format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_task_template#client_token ConnectTaskTemplate#client_token}
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The constraints for the task template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_task_template#constraints ConnectTaskTemplate#constraints}
	Constraints *ConnectTaskTemplateConstraints `field:"optional" json:"constraints" yaml:"constraints"`
	// The identifier of the contact flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_task_template#contact_flow_arn ConnectTaskTemplate#contact_flow_arn}
	ContactFlowArn *string `field:"optional" json:"contactFlowArn" yaml:"contactFlowArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_task_template#defaults ConnectTaskTemplate#defaults}.
	Defaults interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// The description of the task template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_task_template#description ConnectTaskTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The list of task template's fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_task_template#fields ConnectTaskTemplate#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The name of the task template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_task_template#name ConnectTaskTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The identifier of the contact flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_task_template#self_assign_contact_flow_arn ConnectTaskTemplate#self_assign_contact_flow_arn}
	SelfAssignContactFlowArn *string `field:"optional" json:"selfAssignContactFlowArn" yaml:"selfAssignContactFlowArn"`
	// The status of the task template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_task_template#status ConnectTaskTemplate#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_task_template#tags ConnectTaskTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

