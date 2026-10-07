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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/agentregistry_registry#name AgentregistryRegistry#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration for the registry's record approval workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/agentregistry_registry#approval_configuration AgentregistryRegistry#approval_configuration}
	ApprovalConfiguration *AgentregistryRegistryApprovalConfiguration `field:"optional" json:"approvalConfiguration" yaml:"approvalConfiguration"`
	// The type of authorizer that controls how consumers access the registry's search and MCP invoke operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/agentregistry_registry#authorizer_type AgentregistryRegistry#authorizer_type}
	AuthorizerType *string `field:"optional" json:"authorizerType" yaml:"authorizerType"`
	// Specifies whether auto-detection is requested for the registry.
	//
	// Must be specified together with AutoDetectionScope. Setting this to true is necessary but not sufficient for auto-detection to become active; the preconditions of the configured scope must also be met. To turn auto-detection off, explicitly set this to false - removing AutoDetectionEnabled and AutoDetectionScope from the template is a no-op and leaves the existing auto-detection settings unchanged. A registry cannot be deleted while auto-detection is enabled: set this to false and update the stack before deleting the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/agentregistry_registry#auto_detection_enabled AgentregistryRegistry#auto_detection_enabled}
	AutoDetectionEnabled interface{} `field:"optional" json:"autoDetectionEnabled" yaml:"autoDetectionEnabled"`
	// The source from which resources are detected. ORGANIZATION sources resources from all member accounts of an AWS Organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/agentregistry_registry#auto_detection_scope AgentregistryRegistry#auto_detection_scope}
	AutoDetectionScope *string `field:"optional" json:"autoDetectionScope" yaml:"autoDetectionScope"`
	// The description of the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/agentregistry_registry#description AgentregistryRegistry#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Discovery configuration for the registry. Controls how consumers are authorized to search the registry and invoke its MCP endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/agentregistry_registry#discovery_configuration AgentregistryRegistry#discovery_configuration}
	DiscoveryConfiguration *AgentregistryRegistryDiscoveryConfiguration `field:"optional" json:"discoveryConfiguration" yaml:"discoveryConfiguration"`
	// The server-side encryption configuration for a registry.
	//
	// Specifies a customer managed key used to encrypt the registry's content. When omitted, the registry's content is encrypted with an AWS owned key. You cannot change the encryption configuration after registry creation. Specifying a different KMS key, adding this property to an existing registry, or removing it replaces the registry: CloudFormation creates a new registry with a new Amazon Resource Name (ARN) and then deletes the original, including all registry records it contains. Registry records that are not managed by the stack are not re-created in the new registry, and if any remain in the original registry its deletion fails and it is left behind.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/agentregistry_registry#encryption_configuration AgentregistryRegistry#encryption_configuration}
	EncryptionConfiguration *AgentregistryRegistryEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Tags to assign to the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/agentregistry_registry#tags AgentregistryRegistry#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

