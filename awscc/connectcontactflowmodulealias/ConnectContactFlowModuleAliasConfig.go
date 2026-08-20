// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcontactflowmodulealias

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectContactFlowModuleAliasConfig struct {
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
	// The identifier of the contact flow module (ARN) this alias is tied to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_contact_flow_module_alias#contact_flow_module_id ConnectContactFlowModuleAlias#contact_flow_module_id}
	ContactFlowModuleId *string `field:"required" json:"contactFlowModuleId" yaml:"contactFlowModuleId"`
	// The version number of the contact flow module this alias points to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_contact_flow_module_alias#contact_flow_module_version ConnectContactFlowModuleAlias#contact_flow_module_version}
	ContactFlowModuleVersion *float64 `field:"required" json:"contactFlowModuleVersion" yaml:"contactFlowModuleVersion"`
	// The name of the alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_contact_flow_module_alias#name ConnectContactFlowModuleAlias#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_contact_flow_module_alias#description ConnectContactFlowModuleAlias#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

