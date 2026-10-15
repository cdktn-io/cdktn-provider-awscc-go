// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistry


type AgentregistryRegistryCustomMetadataSchemaConfigurationRecordTypeSchemaOverrides struct {
	// The record type that this schema override applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/agentregistry_registry#record_type AgentregistryRegistry#record_type}
	RecordType *string `field:"optional" json:"recordType" yaml:"recordType"`
	// The JSON Schema for the specified record type. Must follow the same structural rules as the default schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/agentregistry_registry#schema AgentregistryRegistry#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

