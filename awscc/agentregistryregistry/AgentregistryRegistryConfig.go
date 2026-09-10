// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistry

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentregistryRegistryConfig struct {
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
	// The name of the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/agentregistry_registry#name AgentregistryRegistry#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration for the registry's record approval workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/agentregistry_registry#approval_configuration AgentregistryRegistry#approval_configuration}
	ApprovalConfiguration *AgentregistryRegistryApprovalConfiguration `field:"optional" json:"approvalConfiguration" yaml:"approvalConfiguration"`
	// The type of authorizer that controls how consumers access the registry's search and MCP invoke operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/agentregistry_registry#authorizer_type AgentregistryRegistry#authorizer_type}
	AuthorizerType *string `field:"optional" json:"authorizerType" yaml:"authorizerType"`
	// The description of the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/agentregistry_registry#description AgentregistryRegistry#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Discovery configuration for the registry. Controls how consumers are authorized to search the registry and invoke its MCP endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/agentregistry_registry#discovery_configuration AgentregistryRegistry#discovery_configuration}
	DiscoveryConfiguration *AgentregistryRegistryDiscoveryConfiguration `field:"optional" json:"discoveryConfiguration" yaml:"discoveryConfiguration"`
	// Tags to assign to the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/agentregistry_registry#tags AgentregistryRegistry#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

