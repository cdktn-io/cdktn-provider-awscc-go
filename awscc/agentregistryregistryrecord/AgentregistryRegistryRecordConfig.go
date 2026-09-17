// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentregistryRegistryRecordConfig struct {
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
	// The typed set of descriptors for a registry record.
	//
	// Exactly one descriptor field is populated based on the record type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#descriptors AgentregistryRegistryRecord#descriptors}
	Descriptors *AgentregistryRegistryRecordDescriptors `field:"required" json:"descriptors" yaml:"descriptors"`
	// The name of the registry record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#name AgentregistryRegistryRecord#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the registry record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#record_type AgentregistryRegistryRecord#record_type}
	RecordType *string `field:"required" json:"recordType" yaml:"recordType"`
	// The description of the registry record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#description AgentregistryRegistryRecord#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The human-readable display name of the registry record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#display_name AgentregistryRegistryRecord#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The version of the registry record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#record_version AgentregistryRegistryRecord#record_version}
	RecordVersion *string `field:"optional" json:"recordVersion" yaml:"recordVersion"`
	// The identifier of the registry in which to create the record.
	//
	// You can specify either the registry ID or the registry Amazon Resource Name (ARN). Use the ARN form to reference a registry shared from another account via AWS Resource Access Manager (RAM).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#registry_id AgentregistryRegistryRecord#registry_id}
	RegistryId *string `field:"optional" json:"registryId" yaml:"registryId"`
	// Tags to assign to the registry record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#tags AgentregistryRegistryRecord#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

